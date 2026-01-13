//go:build linux && gofuzz

package userns

func FuzzUIDMap(uidmap []byte) int {
	_ = uidMapInUserNS(string(uidmap))
	return 1
}
// ID-1768294467-a086a773
