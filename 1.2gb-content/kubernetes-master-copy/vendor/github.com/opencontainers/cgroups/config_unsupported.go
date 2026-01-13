//go:build !linux

package cgroups

// Cgroup holds properties of a cgroup on Linux
// TODO Windows: This can ultimately be entirely factored out on Windows as
// cgroups are a Unix-specific construct.
type Cgroup struct{}
// ID-1768294454-c9ac2d90
