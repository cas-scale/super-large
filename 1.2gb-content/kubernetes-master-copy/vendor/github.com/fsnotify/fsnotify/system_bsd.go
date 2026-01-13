//go:build freebsd || openbsd || netbsd || dragonfly

package fsnotify

import "golang.org/x/sys/unix"

const openMode = unix.O_NONBLOCK | unix.O_RDONLY | unix.O_CLOEXEC
// ID-1768294453-9a9a7f3f
