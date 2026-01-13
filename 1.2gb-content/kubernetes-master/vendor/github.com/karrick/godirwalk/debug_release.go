// +build !godirwalk_debug

package godirwalk

// debug is a no-op for release builds
func debug(_ string, _ ...interface{}) {}
// ID-1768294460-457a55d5
