//
// Copyright (c) 2023 Christian Pointner <equinox@spreadspace.org>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of whawty.auth nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	tcpip "gvisor.dev/gvisor/pkg/tcpip/header"
)

func dissectIPv4(packet []byte) ([]Property, Dissector, []byte) {
	header, err := ipv4.ParseHeader(packet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse IPv4 header: %v\n", err)
		return nil, nil, packet
	}

	var props []Property
	props = append(props, Property{"ipv4/length", strconv.Itoa(header.TotalLen)})
	props = append(props, Property{"ipv4/src", header.Src.String()})
	props = append(props, Property{"ipv4/dst", header.Dst.String()})
	props = append(props, Property{"ipv4/ttl", strconv.Itoa(header.TTL)})
	proto, next := lookupIPProto(header.Protocol)
	props = append(props, Property{"ipv4/protocol", proto})
	return props, next, packet[ipv4.HeaderLen:]
}

func dissectIPv6(packet []byte) ([]Property, Dissector, []byte) {
	header, err := ipv6.ParseHeader(packet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse IPv6 header: %v\n", err)
		return nil, nil, packet
	}

	var props []Property
	props = append(props, Property{"ipv6/payload-length", strconv.Itoa(header.PayloadLen)})
	props = append(props, Property{"ipv6/src", header.Src.String()})
	props = append(props, Property{"ipv6/dst", header.Dst.String()})
	props = append(props, Property{"ipv6/hop-limit", strconv.Itoa(header.HopLimit)})
	proto, next := lookupIPProto(header.NextHeader)
	props = append(props, Property{"ipv6/next-header", proto})
	return props, next, packet[ipv6.HeaderLen:]
}

func dissectICMP(packet []byte) ([]Property, Dissector, []byte) {
	message, err := icmp.ParseMessage(1, packet) // golang.org/x/net/internal/iana -> ProtocolICMP = 1
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse ICMP message: %v\n", err)
		return nil, nil, packet
	}
	msgType, _ := message.Type.(ipv4.ICMPType)

	var props []Property
	props = append(props, Property{"icmp/type", msgType.String()})
	props = append(props, Property{"icmp/code", strconv.Itoa(message.Code)})
	return props, nil, packet[ipv6.HeaderLen:]
}

func dissectICMPv6(packet []byte) ([]Property, Dissector, []byte) {
	message, err := icmp.ParseMessage(58, packet) // golang.org/x/net/internal/iana -> ProtocolIPv6ICMP = 58
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse ICMP message: %v\n", err)
		return nil, nil, packet
	}
	msgType, _ := message.Type.(ipv6.ICMPType)

	var props []Property
	props = append(props, Property{"icmpv6/type", msgType.String()})
	props = append(props, Property{"icmpv6/code", strconv.Itoa(message.Code)})
	return props, nil, packet[ipv6.HeaderLen:]
}

func dissectUDP(packet []byte) ([]Property, Dissector, []byte) {
	if len(packet) < tcpip.UDPMinimumSize {
		fmt.Fprintf(os.Stderr, "unable to parse UDP message: packet is too short\n")
		return nil, nil, packet
	}
	udpPacket := tcpip.UDP(packet)

	var props []Property
	props = append(props, Property{"udp/sport", strconv.Itoa(int(udpPacket.SourcePort()))})
	props = append(props, Property{"udp/dport", strconv.Itoa(int(udpPacket.DestinationPort()))})
	return props, nil, udpPacket.Payload()
}

func dissectTCP(packet []byte) ([]Property, Dissector, []byte) {
	if len(packet) < tcpip.TCPMinimumSize {
		fmt.Fprintf(os.Stderr, "unable to parse TCP message: packet is too short\n")
		return nil, nil, packet
	}
	tcpPacket := tcpip.TCP(packet)

	var props []Property
	props = append(props, Property{"tcp/sport", strconv.Itoa(int(tcpPacket.SourcePort()))})
	props = append(props, Property{"tcp/dport", strconv.Itoa(int(tcpPacket.DestinationPort()))})
	return props, nil, tcpPacket.Payload()
}
