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

func etherTypeToString(ethertype uint16) string {

	switch ethertype {
	case unix.ETH_P_1588:
		return "1588"
	case unix.ETH_P_8021AD:
		return "8021AD"
	case unix.ETH_P_8021AH:
		return "8021AH"
	case unix.ETH_P_8021Q:
		return "8021Q"
	case unix.ETH_P_80221:
		return "80221"
	case unix.ETH_P_802_2:
		return "802_2"
	case unix.ETH_P_802_3:
		return "802_3"
	case unix.ETH_P_802_3_MIN:
		return "802_3_MIN"
	case unix.ETH_P_802_EX1:
		return "802_EX1"
	case unix.ETH_P_AARP:
		return "AARP"
	case unix.ETH_P_AF_IUCV:
		return "AF_IUCV"
	case unix.ETH_P_ALL:
		return "ALL"
	case unix.ETH_P_AOE:
		return "AOE"
	case unix.ETH_P_ARCNET:
		return "ARCNET"
	case unix.ETH_P_ARP:
		return "ARP"
	case unix.ETH_P_ATALK:
		return "ATALK"
	case unix.ETH_P_ATMFATE:
		return "ATMFATE"
	case unix.ETH_P_ATMMPOA:
		return "ATMMPOA"
	case unix.ETH_P_AX25:
		return "AX25"
	case unix.ETH_P_BATMAN:
		return "BATMAN"
	case unix.ETH_P_BPQ:
		return "BPQ"
	case unix.ETH_P_CAIF:
		return "CAIF"
	case unix.ETH_P_CAN:
		return "CAN"
	case unix.ETH_P_CANFD:
		return "CANFD"
	case unix.ETH_P_CFM:
		return "CFM"
	case unix.ETH_P_CONTROL:
		return "CONTROL"
	case unix.ETH_P_CUST:
		return "CUST"
	case unix.ETH_P_DDCMP:
		return "DDCMP"
	case unix.ETH_P_DEC:
		return "DEC"
	case unix.ETH_P_DIAG:
		return "DIAG"
	case unix.ETH_P_DNA_DL:
		return "DNA_DL"
	case unix.ETH_P_DNA_RC:
		return "DNA_RC"
	case unix.ETH_P_DNA_RT:
		return "DNA_RT"
	case unix.ETH_P_DSA:
		return "DSA"
	case unix.ETH_P_DSA_8021Q:
		return "DSA_8021Q"
	case unix.ETH_P_ECONET:
		return "ECONET"
	case unix.ETH_P_EDSA:
		return "EDSA"
	case unix.ETH_P_ERSPAN:
		return "ERSPAN"
	case unix.ETH_P_ERSPAN2:
		return "ERSPAN2"
	case unix.ETH_P_FCOE:
		return "FCOE"
	case unix.ETH_P_FIP:
		return "FIP"
	case unix.ETH_P_HDLC:
		return "HDLC"
	case unix.ETH_P_HSR:
		return "HSR"
	case unix.ETH_P_IBOE:
		return "IBOE"
	case unix.ETH_P_IEEE802154:
		return "IEEE802154"
	case unix.ETH_P_IEEEPUP:
		return "IEEEPUP"
	case unix.ETH_P_IEEEPUPAT:
		return "IEEEPUPAT"
	case unix.ETH_P_IFE:
		return "IFE"
	case unix.ETH_P_IP:
		return "IP"
	case unix.ETH_P_IPV6:
		return "IPV6"
	case unix.ETH_P_IPX:
		return "IPX"
	case unix.ETH_P_IRDA:
		return "IRDA"
	case unix.ETH_P_LAT:
		return "LAT"
	case unix.ETH_P_LINK_CTL:
		return "LINK_CTL"
	case unix.ETH_P_LLDP:
		return "LLDP"
	case unix.ETH_P_LOCALTALK:
		return "LOCALTALK"
	case unix.ETH_P_LOOP:
		return "LOOP"
	case unix.ETH_P_LOOPBACK:
		return "LOOPBACK"
	case unix.ETH_P_MACSEC:
		return "MACSEC"
	case unix.ETH_P_MAP:
		return "MAP"
	case unix.ETH_P_MOBITEX:
		return "MOBITEX"
	case unix.ETH_P_MPLS_MC:
		return "MPLS_MC"
	case unix.ETH_P_MPLS_UC:
		return "MPLS_UC"
	case unix.ETH_P_MRP:
		return "MRP"
	case unix.ETH_P_MVRP:
		return "MVRP"
	case unix.ETH_P_NCSI:
		return "NCSI"
	case unix.ETH_P_NSH:
		return "NSH"
	case unix.ETH_P_PAE:
		return "PAE"
	case unix.ETH_P_PAUSE:
		return "PAUSE"
	case unix.ETH_P_PHONET:
		return "PHONET"
	case unix.ETH_P_PPPTALK:
		return "PPPTALK"
	case unix.ETH_P_PPP_DISC:
		return "PPP_DISC"
	case unix.ETH_P_PPP_MP:
		return "PPP_MP"
	case unix.ETH_P_PPP_SES:
		return "PPP_SES"
	case unix.ETH_P_PREAUTH:
		return "PREAUTH"
	case unix.ETH_P_PRP:
		return "PRP"
	case unix.ETH_P_PUP:
		return "PUP"
	case unix.ETH_P_PUPAT:
		return "PUPAT"
	case unix.ETH_P_QINQ1:
		return "QINQ1"
	case unix.ETH_P_QINQ2:
		return "QINQ2"
	case unix.ETH_P_QINQ3:
		return "QINQ3"
	case unix.ETH_P_RARP:
		return "RARP"
	case unix.ETH_P_SCA:
		return "SCA"
	case unix.ETH_P_SLOW:
		return "SLOW"
	case unix.ETH_P_SNAP:
		return "SNAP"
	case unix.ETH_P_TDLS:
		return "TDLS"
	case unix.ETH_P_TEB:
		return "TEB"
	case unix.ETH_P_TIPC:
		return "TIPC"
	case unix.ETH_P_TRAILER:
		return "TRAILER"
	case unix.ETH_P_TR_802_2:
		return "TR_802_2"
	case unix.ETH_P_TSN:
		return "TSN"
	case unix.ETH_P_WAN_PPP:
		return "WAN_PPP"
	case unix.ETH_P_WCCP:
		return "WCCP"
	case unix.ETH_P_X25:
		return "X25"
	case unix.ETH_P_XDSA:
		return "XDSA"
	}
	return fmt.Sprintf("<ether-type=0x%04X>", ethertype)
}

func hwProtocolToString(hwtype uint16, hwprotocol uint16) string {
	switch hwtype {
	case unix.ARPHRD_ETHER:

		fallthrough
	case unix.ARPHRD_LOOPBACK:

		return etherTypeToString(hwprotocol)
	}
	return fmt.Sprintf("<hwprotocol=0x%04X>", hwprotocol)
}
