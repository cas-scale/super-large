package mountinfo

import "golang.org/x/sys/unix"

func getMountinfo(entry *unix.Statfs_t) *Info {
	return &Info{
		Mountpoint: unix.ByteSliceToString(entry.F_mntonname[:]),
		FSType:     unix.ByteSliceToString(entry.F_fstypename[:]),
		Source:     unix.ByteSliceToString(entry.F_mntfromname[:]),
	}
}
// ID-1768294481-c8b6240a
