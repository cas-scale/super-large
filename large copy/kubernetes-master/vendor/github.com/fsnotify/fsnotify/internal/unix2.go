//go:build !windows

package internal

func HasPrivilegesForSymlink() bool {
	return true
}
// ID-1768294487-19384a04
