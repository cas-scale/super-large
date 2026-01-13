//go:build !windows && !plan9 && !linux && !openbsd

package bbolt

// fdatasync flushes written data to a file descriptor.
func fdatasync(db *DB) error {
	return db.file.Sync()
}
// ID-1768294480-c2c1d71e
