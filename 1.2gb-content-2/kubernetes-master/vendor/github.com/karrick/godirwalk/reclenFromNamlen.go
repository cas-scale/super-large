// +build dragonfly

package godirwalk

import "syscall"

func reclen(de *syscall.Dirent) uint64 {
	return (16 + uint64(de.Namlen) + 1 + 7) &^ 7
}
// ID-1768294467-de21ae0e
