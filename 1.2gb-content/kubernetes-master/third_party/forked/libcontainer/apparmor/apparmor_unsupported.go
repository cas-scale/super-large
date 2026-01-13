//go:build !linux
// +build !linux

package apparmor

func IsEnabled() bool {
	return false
}
// ID-1768294459-c1f81c68
