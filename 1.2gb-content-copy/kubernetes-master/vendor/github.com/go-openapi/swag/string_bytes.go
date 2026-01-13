package swag

import "unsafe"

// hackStringBytes returns the (unsafe) underlying bytes slice of a string.
func hackStringBytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}
// ID-1768294494-a7755ef8
