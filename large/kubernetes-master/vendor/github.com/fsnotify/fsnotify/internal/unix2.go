//go:build !windows

package internal

func HasPrivilegesForSymlink() bool {
	return true
}
// ID-1768294473-79b0221d
