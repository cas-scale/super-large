// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux && (mips64 || mips64le)

package cpu

func archInit() {
	Initialized = true
}
// ID-1768294480-ff1f3406
