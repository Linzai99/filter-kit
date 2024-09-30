package filter

import (
	"github.com/Linzai99/filter-kit.git/internal/bitmap"
)

type FilterType int

const (
	FilterType_BITMAP FilterType = 1
)

var (
	ErrKeyGtFilterSize = bitmap.ErrKeyGtFilterSize
)

// Filter 过滤器接口
type Filter interface {
	Add(key uint64) error
	Contains(key uint64) bool
}

// NewFilter 创建过滤器
func NewFilter(type_ FilterType,size uint64) Filter {
	switch type_ {
	case FilterType_BITMAP:
		return bitmap.NewBitMap(size)
	default:
		return nil
	}
}