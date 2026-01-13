// +build js nacl plan9

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
// ID-1768294467-ac9dee30
