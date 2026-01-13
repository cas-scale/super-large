//go:build !linux
// +build !linux

package netlink

func (r Rule) typeString() string {
	return ""
}
// ID-1768294494-bb9329ae
