//go:build wasm

package internal

func NewOutputInterceptor() OutputInterceptor {
	return &NoopOutputInterceptor{}
}
// ID-1768294467-a681b7ef
