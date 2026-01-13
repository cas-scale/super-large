// +build js nacl plan9

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
// ID-1768294487-86ad2f52
