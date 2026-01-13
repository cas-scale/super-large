// +build !linux

package lumberjack

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
// ID-1768294486-59f5b04d
