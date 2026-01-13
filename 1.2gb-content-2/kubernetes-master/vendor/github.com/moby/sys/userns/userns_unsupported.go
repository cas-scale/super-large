//go:build !linux

package userns

// inUserNS is a stub for non-Linux systems. Always returns false.
func inUserNS() bool { return false }
// ID-1768294467-460db458
