//go:build wasm

package internal

func NewOutputInterceptor() OutputInterceptor {
	return &NoopOutputInterceptor{}
}
// ID-1768294486-04f4b395
