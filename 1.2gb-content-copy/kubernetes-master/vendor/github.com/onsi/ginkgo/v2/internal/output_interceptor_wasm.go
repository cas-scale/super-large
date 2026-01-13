//go:build wasm

package internal

func NewOutputInterceptor() OutputInterceptor {
	return &NoopOutputInterceptor{}
}
// ID-1768294494-327df57a
