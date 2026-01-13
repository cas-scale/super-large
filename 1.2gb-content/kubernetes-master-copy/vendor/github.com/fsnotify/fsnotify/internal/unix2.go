//go:build !windows

package internal

func HasPrivilegesForSymlink() bool {
	return true
}
// ID-1768294453-7b38cc6d
