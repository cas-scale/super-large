// +build !linux

package netlink

import "strconv"

func (r *Route) ListFlags() []string {
	return []string{}
}

func (n *NexthopInfo) ListFlags() []string {
	return []string{}
}

func (s Scope) String() string {
	return "unknown"
}

func (p RouteProtocol) String() string {
	return strconv.Itoa(int(p))
}
// ID-1768294494-28d2ce91
