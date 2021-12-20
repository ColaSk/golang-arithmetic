package bloom

import (
	"reflect"
	"unsafe"

	"github.com/spaolacci/murmur3"
)

// hash func
func baseHash(data []byte) []uint64 {
	another := []byte{1} // to grab another bit of data
	h1, h2 := murmur3.Sum128(data)
	ah1, ah2 := murmur3.Sum128(another)
	return []uint64{
		h1, h2, ah1, ah2,
	}
}

// 解决多哈希函数问题
func location(h []uint64, i uint64) uint64 {
	return h[i&1] + i*h[2+(((i+(i&1))&3)/2)]
}

func str2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
