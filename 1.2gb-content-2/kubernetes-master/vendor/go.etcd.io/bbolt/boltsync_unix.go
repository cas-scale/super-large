//go:build !windows && !plan9 && !linux && !openbsd

package bbolt

// fdatasync flushes written data to a file descriptor.
func fdatasync(db *DB) error {
	return db.file.Sync()
}
// ID-1768294466-8f6e3bdf
