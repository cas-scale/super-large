// +build appengine

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
// ID-1768294474-15499092
