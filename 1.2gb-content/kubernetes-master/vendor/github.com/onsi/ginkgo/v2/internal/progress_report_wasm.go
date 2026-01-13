//go:build wasm

package internal

import (
	"os"
	"syscall"
)

var PROGRESS_SIGNALS = []os.Signal{syscall.SIGUSR1}
// ID-1768294460-8f0b71f4
