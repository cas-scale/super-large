// +build !linux

package lumberjack

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
// ID-1768294480-2f74872a
