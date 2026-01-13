//go:build !linux
// +build !linux

package netlink

func (r Rule) typeString() string {
	return ""
}
// ID-1768294480-5d3128ee
