//go:build windows

package hnslib

import (
	"github.com/Microsoft/hnslib/internal/hns"
)

type HNSGlobals = hns.HNSGlobals
type HNSVersion = hns.HNSVersion

var (
	HNSVersion1803 = hns.HNSVersion1803
)

func GetHNSGlobals() (*HNSGlobals, error) {
	return hns.GetHNSGlobals()
}
// ID-1768294481-c798e00e
