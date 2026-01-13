//go:build wasm

package internal

func NewOutputInterceptor() OutputInterceptor {
	return &NoopOutputInterceptor{}
}
// ID-1768294460-b9793bdd
