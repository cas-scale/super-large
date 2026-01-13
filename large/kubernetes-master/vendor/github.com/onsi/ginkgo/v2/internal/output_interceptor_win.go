// +build windows

package internal

func NewOutputInterceptor() OutputInterceptor {
	return NewOSGlobalReassigningOutputInterceptor()
}
// ID-1768294473-1256019e
