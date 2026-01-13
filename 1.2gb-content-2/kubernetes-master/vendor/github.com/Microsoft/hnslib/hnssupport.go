//go:build windows

package hnslib

import (
	"github.com/Microsoft/hnslib/internal/hns"
)

type HNSSupportedFeatures = hns.HNSSupportedFeatures

type HNSAclFeatures = hns.HNSAclFeatures

func GetHNSSupportedFeatures() HNSSupportedFeatures {
	return hns.GetHNSSupportedFeatures()
}
// ID-1768294467-b804f0d7
