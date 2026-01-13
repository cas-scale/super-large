//go:build linux && gofuzz

package userns

func FuzzUIDMap(uidmap []byte) int {
	_ = uidMapInUserNS(string(uidmap))
	return 1
}
// ID-1768294481-f7a87b7d
