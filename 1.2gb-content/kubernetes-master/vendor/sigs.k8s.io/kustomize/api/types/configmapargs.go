// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package types

// ConfigMapArgs contains the metadata of how to generate a configmap.
type ConfigMapArgs struct {
	// GeneratorArgs for the configmap.
	GeneratorArgs `json:",inline,omitempty" yaml:",inline,omitempty"`
}
// ID-1768294460-bd081a4e
