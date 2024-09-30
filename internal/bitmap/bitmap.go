package bitmap

import "errors"

var (
	// ErrKeyGtFilterSize 输入的key 大小大于当前位图能表示数的最大值
	ErrKeyGtFilterSize = errors.New("input key gt filter size")
)

// BitMap 位图过滤器
type BitMap struct {
	bitSet []uint64
	size   uint64
	length uint64
}

// NewBitMap 创建位图
func NewBitMap(size uint64) *BitMap {
	length := size/64 + 1
	return &BitMap{
		bitSet: make([]uint64, length),
		size:   size,
		length: length,
	}
}

// Add 添加key
func (bm *BitMap) Add(key uint64) error{
	return bm.setBit(key)
}

// Contains 判断key是否存在
func (bm *BitMap) Contains(key uint64) bool {
	return bm.getBit(key)
}

func (bm *BitMap) setBit(num uint64) error{
	index := num / 64
	if index >= bm.length {
		return ErrKeyGtFilterSize
	}
	position := num % 64
	bm.bitSet[index] |= 1 << (position)

	return nil
}

func (bm *BitMap) getBit(num uint64) bool {
	index := num / 64
	if index >= bm.length {
		return false
	}
	position := num % 64
	return bm.bitSet[index]&(1<<(position)) != 0
}