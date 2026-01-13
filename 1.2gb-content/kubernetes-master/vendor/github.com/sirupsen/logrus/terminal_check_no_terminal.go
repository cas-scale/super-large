// +build js nacl plan9

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
// ID-1768294460-7e1666bd
