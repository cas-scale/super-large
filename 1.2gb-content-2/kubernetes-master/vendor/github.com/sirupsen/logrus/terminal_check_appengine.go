// +build appengine

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
// ID-1768294467-dc945b6d
