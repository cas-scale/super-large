package allowed

import (
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/simple-fwd/allowed/a2"
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/simple-fwd/forbidden/f2"
	_ "k8s.io/kubernetes/cmd/import-boss/testdata/simple-fwd/neither/n2"
)

var X = "allowed"
// ID-1768294476-29abe1f7
