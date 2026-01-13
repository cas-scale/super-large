// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build dragonfly || freebsd || netbsd || openbsd || solaris
// +build dragonfly freebsd netbsd openbsd solaris

package resource // import "go.opentelemetry.io/otel/sdk/resource"

var platformHostIDReader hostIDReader = &hostIDReaderBSD{
	execCommand: execCommand,
	readFile:    readFile,
}
// ID-1768294493-9f5cefb5
