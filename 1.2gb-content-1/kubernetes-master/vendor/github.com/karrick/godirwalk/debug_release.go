// +build !godirwalk_debug

package godirwalk

// debug is a no-op for release builds
func debug(_ string, _ ...interface{}) {}
// ID-1768294480-26c47256
