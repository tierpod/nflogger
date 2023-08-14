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

func hwTypeToString(hwtype uint16) string {
	switch hwtype {
	case unix.ARPHRD_6LOWPAN:
		return "6LOWPAN"
	case unix.ARPHRD_ADAPT:
		return "ADAPT"
	case unix.ARPHRD_APPLETLK:
		return "APPLETLK"
	case unix.ARPHRD_ARCNET:
		return "ARCNET"
	case unix.ARPHRD_ASH:
		return "ASH"
	case unix.ARPHRD_ATM:
		return "ATM"
	case unix.ARPHRD_AX25:
		return "AX25"
	case unix.ARPHRD_BIF:
		return "BIF"
	case unix.ARPHRD_CAIF:
		return "CAIF"
	case unix.ARPHRD_CAN:
		return "CAN"
	case unix.ARPHRD_CHAOS:
		return "CHAOS"
	case unix.ARPHRD_CSLIP:
		return "CSLIP"
	case unix.ARPHRD_CSLIP6:
		return "CSLIP6"
	case unix.ARPHRD_DDCMP:
		return "DDCMP"
	case unix.ARPHRD_DLCI:
		return "DLCI"
	case unix.ARPHRD_ECONET:
		return "ECONET"
	case unix.ARPHRD_EETHER:
		return "EETHER"
	case unix.ARPHRD_ETHER:
		return "ETHER"
	case unix.ARPHRD_EUI64:
		return "EUI64"
	case unix.ARPHRD_FCAL:
		return "FCAL"
	case unix.ARPHRD_FCFABRIC:
		return "FCFABRIC"
	case unix.ARPHRD_FCPL:
		return "FCPL"
	case unix.ARPHRD_FCPP:
		return "FCPP"
	case unix.ARPHRD_FDDI:
		return "FDDI"
	case unix.ARPHRD_FRAD:
		return "FRAD"
	case unix.ARPHRD_HDLC:
		return "HDLC"
	case unix.ARPHRD_HIPPI:
		return "HIPPI"
	case unix.ARPHRD_HWX25:
		return "HWX25"
	case unix.ARPHRD_IEEE1394:
		return "IEEE1394"
	case unix.ARPHRD_IEEE802:
		return "IEEE802"
	case unix.ARPHRD_IEEE80211:
		return "IEEE80211"
	case unix.ARPHRD_IEEE80211_PRISM:
		return "IEEE80211_PRISM"
	case unix.ARPHRD_IEEE80211_RADIOTAP:
		return "IEEE80211_RADIOTAP"
	case unix.ARPHRD_IEEE802154:
		return "IEEE802154"
	case unix.ARPHRD_IEEE802154_MONITOR:
		return "IEEE802154_MONITOR"
	case unix.ARPHRD_IEEE802_TR:
		return "IEEE802_TR"
	case unix.ARPHRD_INFINIBAND:
		return "INFINIBAND"
	case unix.ARPHRD_IP6GRE:
		return "IP6GRE"
	case unix.ARPHRD_IPDDP:
		return "IPDDP"
	case unix.ARPHRD_IPGRE:
		return "IPGRE"
	case unix.ARPHRD_IRDA:
		return "IRDA"
	case unix.ARPHRD_LAPB:
		return "LAPB"
	case unix.ARPHRD_LOCALTLK:
		return "LOCALTLK"
	case unix.ARPHRD_LOOPBACK:
		return "LOOPBACK"
	case unix.ARPHRD_METRICOM:
		return "METRICOM"
	case unix.ARPHRD_NETLINK:
		return "NETLINK"
	case unix.ARPHRD_NETROM:
		return "NETROM"
	case unix.ARPHRD_NONE:
		return "NONE"
	case unix.ARPHRD_PHONET:
		return "PHONET"
	case unix.ARPHRD_PHONET_PIPE:
		return "PHONET_PIPE"
	case unix.ARPHRD_PIMREG:
		return "PIMREG"
	case unix.ARPHRD_PPP:
		return "PPP"
	case unix.ARPHRD_PRONET:
		return "PRONET"
	case unix.ARPHRD_RAWHDLC:
		return "RAWHDLC"
	case unix.ARPHRD_RAWIP:
		return "RAWIP"
	case unix.ARPHRD_ROSE:
		return "ROSE"
	case unix.ARPHRD_RSRVD:
		return "RSRVD"
	case unix.ARPHRD_SIT:
		return "SIT"
	case unix.ARPHRD_SKIP:
		return "SKIP"
	case unix.ARPHRD_SLIP:
		return "SLIP"
	case unix.ARPHRD_SLIP6:
		return "SLIP6"
	case unix.ARPHRD_TUNNEL:
		return "TUNNEL"
	case unix.ARPHRD_TUNNEL6:
		return "TUNNEL6"
	case unix.ARPHRD_VOID:
		return "VOID"
	case unix.ARPHRD_VSOCKMON:
		return "VSOCKMON"
	case unix.ARPHRD_X25:
		return "X25"
	}
	return fmt.Sprintf("<hwtype=0x%04X>", hwtype)
}

func lookupEtherType(ethertype uint16) (string, Dissector) {
	switch ethertype {
	case unix.ETH_P_1588:
		return "1588", nil
	case unix.ETH_P_8021AD:
		return "8021AD", nil
	case unix.ETH_P_8021AH:
		return "8021AH", nil
	case unix.ETH_P_8021Q:
		return "8021Q", nil
	case unix.ETH_P_80221:
		return "80221", nil
	case unix.ETH_P_802_2:
		return "802_2", nil
	case unix.ETH_P_802_3:
		return "802_3", nil
	case unix.ETH_P_802_3_MIN:
		return "802_3_MIN", nil
	case unix.ETH_P_802_EX1:
		return "802_EX1", nil
	case unix.ETH_P_AARP:
		return "AARP", nil
	case unix.ETH_P_AF_IUCV:
		return "AF_IUCV", nil
	case unix.ETH_P_ALL:
		return "ALL", nil
	case unix.ETH_P_AOE:
		return "AOE", nil
	case unix.ETH_P_ARCNET:
		return "ARCNET", nil
	case unix.ETH_P_ARP:
		return "ARP", nil
	case unix.ETH_P_ATALK:
		return "ATALK", nil
	case unix.ETH_P_ATMFATE:
		return "ATMFATE", nil
	case unix.ETH_P_ATMMPOA:
		return "ATMMPOA", nil
	case unix.ETH_P_AX25:
		return "AX25", nil
	case unix.ETH_P_BATMAN:
		return "BATMAN", nil
	case unix.ETH_P_BPQ:
		return "BPQ", nil
	case unix.ETH_P_CAIF:
		return "CAIF", nil
	case unix.ETH_P_CAN:
		return "CAN", nil
	case unix.ETH_P_CANFD:
		return "CANFD", nil
	case unix.ETH_P_CFM:
		return "CFM", nil
	case unix.ETH_P_CONTROL:
		return "CONTROL", nil
	case unix.ETH_P_CUST:
		return "CUST", nil
	case unix.ETH_P_DDCMP:
		return "DDCMP", nil
	case unix.ETH_P_DEC:
		return "DEC", nil
	case unix.ETH_P_DIAG:
		return "DIAG", nil
	case unix.ETH_P_DNA_DL:
		return "DNA_DL", nil
	case unix.ETH_P_DNA_RC:
		return "DNA_RC", nil
	case unix.ETH_P_DNA_RT:
		return "DNA_RT", nil
	case unix.ETH_P_DSA:
		return "DSA", nil
	case unix.ETH_P_DSA_8021Q:
		return "DSA_8021Q", nil
	case unix.ETH_P_ECONET:
		return "ECONET", nil
	case unix.ETH_P_EDSA:
		return "EDSA", nil
	case unix.ETH_P_ERSPAN:
		return "ERSPAN", nil
	case unix.ETH_P_ERSPAN2:
		return "ERSPAN2", nil
	case unix.ETH_P_FCOE:
		return "FCOE", nil
	case unix.ETH_P_FIP:
		return "FIP", nil
	case unix.ETH_P_HDLC:
		return "HDLC", nil
	case unix.ETH_P_HSR:
		return "HSR", nil
	case unix.ETH_P_IBOE:
		return "IBOE", nil
	case unix.ETH_P_IEEE802154:
		return "IEEE802154", nil
	case unix.ETH_P_IEEEPUP:
		return "IEEEPUP", nil
	case unix.ETH_P_IEEEPUPAT:
		return "IEEEPUPAT", nil
	case unix.ETH_P_IFE:
		return "IFE", nil
	case unix.ETH_P_IP:
		return "IP", dissectIPv4
	case unix.ETH_P_IPV6:
		return "IPV6", nil
	case unix.ETH_P_IPX:
		return "IPX", nil
	case unix.ETH_P_IRDA:
		return "IRDA", nil
	case unix.ETH_P_LAT:
		return "LAT", nil
	case unix.ETH_P_LINK_CTL:
		return "LINK_CTL", nil
	case unix.ETH_P_LLDP:
		return "LLDP", nil
	case unix.ETH_P_LOCALTALK:
		return "LOCALTALK", nil
	case unix.ETH_P_LOOP:
		return "LOOP", nil
	case unix.ETH_P_LOOPBACK:
		return "LOOPBACK", nil
	case unix.ETH_P_MACSEC:
		return "MACSEC", nil
	case unix.ETH_P_MAP:
		return "MAP", nil
	case unix.ETH_P_MOBITEX:
		return "MOBITEX", nil
	case unix.ETH_P_MPLS_MC:
		return "MPLS_MC", nil
	case unix.ETH_P_MPLS_UC:
		return "MPLS_UC", nil
	case unix.ETH_P_MRP:
		return "MRP", nil
	case unix.ETH_P_MVRP:
		return "MVRP", nil
	case unix.ETH_P_NCSI:
		return "NCSI", nil
	case unix.ETH_P_NSH:
		return "NSH", nil
	case unix.ETH_P_PAE:
		return "PAE", nil
	case unix.ETH_P_PAUSE:
		return "PAUSE", nil
	case unix.ETH_P_PHONET:
		return "PHONET", nil
	case unix.ETH_P_PPPTALK:
		return "PPPTALK", nil
	case unix.ETH_P_PPP_DISC:
		return "PPP_DISC", nil
	case unix.ETH_P_PPP_MP:
		return "PPP_MP", nil
	case unix.ETH_P_PPP_SES:
		return "PPP_SES", nil
	case unix.ETH_P_PREAUTH:
		return "PREAUTH", nil
	case unix.ETH_P_PRP:
		return "PRP", nil
	case unix.ETH_P_PUP:
		return "PUP", nil
	case unix.ETH_P_PUPAT:
		return "PUPAT", nil
	case unix.ETH_P_QINQ1:
		return "QINQ1", nil
	case unix.ETH_P_QINQ2:
		return "QINQ2", nil
	case unix.ETH_P_QINQ3:
		return "QINQ3", nil
	case unix.ETH_P_RARP:
		return "RARP", nil
	case unix.ETH_P_SCA:
		return "SCA", nil
	case unix.ETH_P_SLOW:
		return "SLOW", nil
	case unix.ETH_P_SNAP:
		return "SNAP", nil
	case unix.ETH_P_TDLS:
		return "TDLS", nil
	case unix.ETH_P_TEB:
		return "TEB", nil
	case unix.ETH_P_TIPC:
		return "TIPC", nil
	case unix.ETH_P_TRAILER:
		return "TRAILER", nil
	case unix.ETH_P_TR_802_2:
		return "TR_802_2", nil
	case unix.ETH_P_TSN:
		return "TSN", nil
	case unix.ETH_P_WAN_PPP:
		return "WAN_PPP", nil
	case unix.ETH_P_WCCP:
		return "WCCP", nil
	case unix.ETH_P_X25:
		return "X25", nil
	case unix.ETH_P_XDSA:
		return "XDSA", nil
	}
	return fmt.Sprintf("<ether-type=0x%04X>", ethertype), nil
}

func lookupHWProtocol(hwtype uint16, hwprotocol uint16) (string, Dissector) {
	switch hwtype {
	case unix.ARPHRD_ETHER:

		fallthrough
	case unix.ARPHRD_LOOPBACK:

		return lookupEtherType(hwprotocol)
	}
	return fmt.Sprintf("<hwprotocol=0x%04X>", hwprotocol), nil
}

func lookupIPProto(ipproto int) (string, Dissector) {
	switch ipproto {
	case unix.IPPROTO_AH:
		return "AH", nil
	case unix.IPPROTO_BEETPH:
		return "BEETPH", nil
	case unix.IPPROTO_COMP:
		return "COMP", nil
	case unix.IPPROTO_DCCP:
		return "DCCP", nil
	case unix.IPPROTO_DSTOPTS:
		return "DSTOPTS", nil
	case unix.IPPROTO_EGP:
		return "EGP", nil
	case unix.IPPROTO_ENCAP:
		return "ENCAP", nil
	case unix.IPPROTO_ESP:
		return "ESP", nil
	case unix.IPPROTO_ETHERNET:
		return "ETHERNET", nil
	case unix.IPPROTO_FRAGMENT:
		return "FRAGMENT", nil
	case unix.IPPROTO_GRE:
		return "GRE", nil
	case unix.IPPROTO_ICMP:
		return "ICMP", nil
	case unix.IPPROTO_ICMPV6:
		return "ICMPV6", nil
	case unix.IPPROTO_IDP:
		return "IDP", nil
	case unix.IPPROTO_IGMP:
		return "IGMP", nil
	case unix.IPPROTO_IP:
		return "IP", nil
	case unix.IPPROTO_IPIP:
		return "IPIP", nil
	case unix.IPPROTO_IPV6:
		return "IPV6", nil
	case unix.IPPROTO_L2TP:
		return "L2TP", nil
	case unix.IPPROTO_MH:
		return "MH", nil
	case unix.IPPROTO_MPLS:
		return "MPLS", nil
	case unix.IPPROTO_MPTCP:
		return "MPTCP", nil
	case unix.IPPROTO_MTP:
		return "MTP", nil
	case unix.IPPROTO_NONE:
		return "NONE", nil
	case unix.IPPROTO_PIM:
		return "PIM", nil
	case unix.IPPROTO_PUP:
		return "PUP", nil
	case unix.IPPROTO_RAW:
		return "RAW", nil
	case unix.IPPROTO_ROUTING:
		return "ROUTING", nil
	case unix.IPPROTO_RSVP:
		return "RSVP", nil
	case unix.IPPROTO_SCTP:
		return "SCTP", nil
	case unix.IPPROTO_TCP:
		return "TCP", nil
	case unix.IPPROTO_TP:
		return "TP", nil
	case unix.IPPROTO_UDP:
		return "UDP", nil
	case unix.IPPROTO_UDPLITE:
		return "UDPLITE", nil
	}
	return fmt.Sprintf("<ipproto=0x%02X>", ipproto), nil
}
