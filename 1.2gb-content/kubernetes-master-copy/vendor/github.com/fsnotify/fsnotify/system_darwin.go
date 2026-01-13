//go:build darwin

package fsnotify

import "golang.org/x/sys/unix"

// note: this constant is not defined on BSD
const openMode = unix.O_EVTONLY | unix.O_CLOEXEC
// ID-1768294453-5189ba4d
