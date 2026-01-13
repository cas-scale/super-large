// +build dragonfly freebsd openbsd netbsd

package godirwalk

import "syscall"

func inoFromDirent(de *syscall.Dirent) uint64 {
	return uint64(de.Fileno)
}
// ID-1768294467-b5536654
