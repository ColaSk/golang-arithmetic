package bloom

import (
	"math"
)

/*
 * 布隆过滤器的实现
 */

/*
 * size:    过滤器大小
 * num:     插入元素个数
 * log2m:   log_2 of size
 * hashnum: 散列函数个数
 * keys:    过滤器
 */

const (
	mod7       = 1<<3 - 1
	bitPerByte = 8
)

type BloomFilter struct {
	size    uint64
	num     uint64
	log2m   uint64
	hashnum uint64
	keys    []byte
}

// 返回位置
// & (f.m - 1) 取余运算的快捷方式
func (f *BloomFilter) location(h uint64) (uint64, uint64) {
	slot := (h / bitPerByte) & (f.size - 1)
	mod := h & mod7
	return slot, mod
}

func (f *BloomFilter) Add(data []byte) {

	h := baseHash(data)

	for i := uint64(0); i < f.hashnum; i++ {
		loc := location(h, i)
		slot, mod := f.location(loc)
		f.keys[slot] |= 1 << mod
	}
}

func (f *BloomFilter) AddString(s string) {
	data := str2Bytes(s)
	f.Add(data)
	f.num++
}

func (f *BloomFilter) ExistsString(s string) bool {
	data := str2Bytes(s)
	return f.exists(data)
}

func (f *BloomFilter) exists(data []byte) bool {
	h := baseHash(data)
	for i := uint64(0); i < f.hashnum; i++ {
		loc := location(h, i)
		slot, mod := f.location(loc)
		if f.keys[slot]&(1<<mod) == 0 {
			return false
		}
	}
	return true
}

func New(size uint64, hashnum uint64) *BloomFilter {
	/*
	* hashnum: 散列函数个数
	* size: 过滤器大小
	 */

	// 将size拉长为2^n
	log2 := uint64(math.Ceil(math.Log2(float64(size))))

	filter := &BloomFilter{
		size:    1 << log2,
		log2m:   log2,
		hashnum: hashnum,
		keys:    make([]byte, 1<<log2),
	}
	return filter
}
