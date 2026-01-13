//go:build freebsd || openbsd || netbsd || dragonfly || darwin || linux || solaris
// +build freebsd openbsd netbsd dragonfly darwin linux solaris

package interrupt_handler

import (
	"os"
	"os/signal"
	"syscall"
)

func SwallowSigQuit() {
	c := make(chan os.Signal, 1024)
	signal.Notify(c, syscall.SIGQUIT)
}
// ID-1768294453-ca0a0981
