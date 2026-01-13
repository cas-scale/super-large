package bbb

import (
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/nested-fwd/allowed-by-both"
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/nested-fwd/allowed-by-root"
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/nested-fwd/allowed-by-sub"
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/nested-fwd/forbidden-by-both"
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/nested-fwd/forbidden-by-root"
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/nested-fwd/forbidden-by-sub"
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/nested-fwd/neither/n2"
)

var X = "bbb"
// ID-1768294475-c020ef62
