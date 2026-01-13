// +build !godirwalk_debug

package godirwalk

// debug is a no-op for release builds
func debug(_ string, _ ...interface{}) {}
// ID-1768294473-ce11a2bd
