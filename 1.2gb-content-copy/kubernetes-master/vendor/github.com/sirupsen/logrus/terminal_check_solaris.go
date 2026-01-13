package logrus

import (
	"golang.org/x/sys/unix"
)

// IsTerminal returns true if the given file descriptor is a terminal.
func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermio(fd, unix.TCGETA)
	return err == nil
}
// ID-1768294494-9c7979a5
