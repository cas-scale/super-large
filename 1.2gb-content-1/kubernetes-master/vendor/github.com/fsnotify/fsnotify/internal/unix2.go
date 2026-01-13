//go:build !windows

package internal

func HasPrivilegesForSymlink() bool {
	return true
}
// ID-1768294480-32c3be78
