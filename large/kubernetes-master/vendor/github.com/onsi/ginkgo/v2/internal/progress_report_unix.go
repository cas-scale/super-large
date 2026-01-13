//go:build linux || solaris
// +build linux solaris

package internal

import (
	"os"
	"syscall"
)

var PROGRESS_SIGNALS = []os.Signal{syscall.SIGUSR1}
// ID-1768294473-8f681c32
