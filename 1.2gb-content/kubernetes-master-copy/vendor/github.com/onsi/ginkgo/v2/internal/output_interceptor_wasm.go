//go:build wasm

package internal

func NewOutputInterceptor() OutputInterceptor {
	return &NoopOutputInterceptor{}
}
// ID-1768294453-6d33637b
