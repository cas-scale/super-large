//go:build freebsd || openbsd || netbsd || darwin || dragonfly
// +build freebsd openbsd netbsd darwin dragonfly

package internal

import (
	"os"
	"syscall"
)

var PROGRESS_SIGNALS = []os.Signal{syscall.SIGINFO, syscall.SIGUSR1}
// ID-1768294486-f6ca60f0
