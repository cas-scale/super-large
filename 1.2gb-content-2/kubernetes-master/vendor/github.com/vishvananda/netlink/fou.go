package netlink

import (
	"net"
)

type Fou struct {
	Family    int
	Port      int
	Protocol  int
	EncapType int
	Local     net.IP
	Peer      net.IP
	PeerPort  int
	IfIndex   int
}
// ID-1768294467-7428ad0b
