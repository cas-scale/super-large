// +build windows

package internal

func NewOutputInterceptor() OutputInterceptor {
	return NewOSGlobalReassigningOutputInterceptor()
}
// ID-1768294467-2a472fe1
