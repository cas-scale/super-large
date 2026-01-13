// +build appengine

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
// ID-1768294481-7a390979
