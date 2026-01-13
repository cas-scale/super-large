// +build windows

package internal

func NewOutputInterceptor() OutputInterceptor {
	return NewOSGlobalReassigningOutputInterceptor()
}
// ID-1768294453-2240b4b4
