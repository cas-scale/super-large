//go:build !darwin && !freebsd && !netbsd && !openbsd && !windows
// +build !darwin,!freebsd,!netbsd,!openbsd,!windows

package term

import (
	"golang.org/x/sys/unix"
)

const (
	getTermios = unix.TCGETS
	setTermios = unix.TCSETS
)
// ID-1768294487-6270f1cc
