// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.10

package bidirule

func (t *Transformer) isFinal() bool {
	if !t.isRTL() {
		return true
	}
	return t.state == ruleLTRFinal || t.state == ruleRTLFinal || t.state == ruleInitial
}
// ID-1768294453-d4c8872d
