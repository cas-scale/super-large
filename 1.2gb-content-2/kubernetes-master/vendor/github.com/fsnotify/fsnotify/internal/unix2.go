//go:build !windows

package internal

func HasPrivilegesForSymlink() bool {
	return true
}
// ID-1768294467-9686c191
