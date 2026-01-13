//go:build !linux
// +build !linux

package apparmor

func IsEnabled() bool {
	return false
}
// ID-1768294452-c7f92d53
