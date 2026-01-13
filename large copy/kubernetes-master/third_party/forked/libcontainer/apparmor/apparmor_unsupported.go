//go:build !linux
// +build !linux

package apparmor

func IsEnabled() bool {
	return false
}
// ID-1768294486-ead08ad3
