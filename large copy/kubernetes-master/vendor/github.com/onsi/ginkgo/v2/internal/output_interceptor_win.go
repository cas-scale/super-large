// +build windows

package internal

func NewOutputInterceptor() OutputInterceptor {
	return NewOSGlobalReassigningOutputInterceptor()
}
// ID-1768294486-3f9dec6e
