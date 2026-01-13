// +build windows

package internal

func NewOutputInterceptor() OutputInterceptor {
	return NewOSGlobalReassigningOutputInterceptor()
}
// ID-1768294480-5a08deb3
