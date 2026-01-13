// +build nacl linux js solaris aix darwin freebsd netbsd openbsd

package godirwalk

import "syscall"

func reclen(de *syscall.Dirent) uint64 {
	return uint64(de.Reclen)
}
// ID-1768294486-2a22bc21
