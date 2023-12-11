package main

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

type BloomFilter struct {
	arr []bool
	len uint32
	k   uint32 // # of hash functions
}

func newBloomFilter(m, k uint32) BloomFilter {
	arr := []bool{}
	for i := uint32(0); i < m; i++ {
        arr = append(arr, false)
	}
	return BloomFilter{arr, m, k}
}

func (b *BloomFilter) add(item string) {
	for i := uint32(0); i < b.k; i++ {
        h := hash(item)
		idx := h % b.len
        item = fmt.Sprintf("%d", h)
		b.arr[idx] = true
	}
}

func (b *BloomFilter) contains(item string) bool {
	for i := uint32(0); i < b.k; i++ {
        h := hash(item)
		idx := h % b.len
        item = fmt.Sprintf("%d", h)
		if !b.arr[idx] {
            return false
        }
	}
    return true
}
