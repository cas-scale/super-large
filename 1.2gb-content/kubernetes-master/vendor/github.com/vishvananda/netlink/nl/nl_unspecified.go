// +build !linux

package nl

import "encoding/binary"

var SupportedNlFamilies = []int{}

func NativeEndian() binary.ByteOrder {
	return nil
}
// ID-1768294460-83923fef
