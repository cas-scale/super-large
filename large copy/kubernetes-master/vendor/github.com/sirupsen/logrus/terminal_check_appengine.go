// +build appengine

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
// ID-1768294487-078f366b
