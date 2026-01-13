//go:build !linux
// +build !linux

package apparmor

func IsEnabled() bool {
	return false
}
// ID-1768294466-fb19ba0a
