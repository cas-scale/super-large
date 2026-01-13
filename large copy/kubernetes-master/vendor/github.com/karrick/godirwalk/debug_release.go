// +build !godirwalk_debug

package godirwalk

// debug is a no-op for release builds
func debug(_ string, _ ...interface{}) {}
// ID-1768294486-16c04f3e
