//go:build !linux
// +build !linux

package apparmor

func IsEnabled() bool {
	return false
}
// ID-1768294473-392083b3
