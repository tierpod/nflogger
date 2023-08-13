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

	"golang.org/x/sys/unix"
)

func ipProtoToString(ipproto int) string {
	switch ipproto {
	case unix.IPPROTO_AH:
		return "AH"
	case unix.IPPROTO_BEETPH:
		return "BEETPH"
	case unix.IPPROTO_COMP:
		return "COMP"
	case unix.IPPROTO_DCCP:
		return "DCCP"
	case unix.IPPROTO_DSTOPTS:
		return "DSTOPTS"
	case unix.IPPROTO_EGP:
		return "EGP"
	case unix.IPPROTO_ENCAP:
		return "ENCAP"
	case unix.IPPROTO_ESP:
		return "ESP"
	case unix.IPPROTO_ETHERNET:
		return "ETHERNET"
	case unix.IPPROTO_FRAGMENT:
		return "FRAGMENT"
	case unix.IPPROTO_GRE:
		return "GRE"
	case unix.IPPROTO_ICMP:
		return "ICMP"
	case unix.IPPROTO_ICMPV6:
		return "ICMPV6"
	case unix.IPPROTO_IDP:
		return "IDP"
	case unix.IPPROTO_IGMP:
		return "IGMP"
	case unix.IPPROTO_IP:
		return "IP"
	case unix.IPPROTO_IPIP:
		return "IPIP"
	case unix.IPPROTO_IPV6:
		return "IPV6"
	case unix.IPPROTO_L2TP:
		return "L2TP"
	case unix.IPPROTO_MH:
		return "MH"
	case unix.IPPROTO_MPLS:
		return "MPLS"
	case unix.IPPROTO_MPTCP:
		return "MPTCP"
	case unix.IPPROTO_MTP:
		return "MTP"
	case unix.IPPROTO_NONE:
		return "NONE"
	case unix.IPPROTO_PIM:
		return "PIM"
	case unix.IPPROTO_PUP:
		return "PUP"
	case unix.IPPROTO_RAW:
		return "RAW"
	case unix.IPPROTO_ROUTING:
		return "ROUTING"
	case unix.IPPROTO_RSVP:
		return "RSVP"
	case unix.IPPROTO_SCTP:
		return "SCTP"
	case unix.IPPROTO_TCP:
		return "TCP"
	case unix.IPPROTO_TP:
		return "TP"
	case unix.IPPROTO_UDP:
		return "UDP"
	case unix.IPPROTO_UDPLITE:
		return "UDPLITE"
	}
	return fmt.Sprintf("<ipproto=0x%02X>", ipproto)
}
