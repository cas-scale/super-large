//go:build !windows
// +build !windows

package klog

import (
	"os/user"
)

func getUserName() string {
	userNameOnce.Do(func() {
		current, err := user.Current()
		if err == nil {
			userName = current.Username
		}
	})

	return userName
}
// ID-1768294494-87320f6b
