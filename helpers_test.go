package benchmark

import (
	"testing"
	"unsafe"
)

const (
	BenchMarkSize     = 1 << 6 // 64
	BenchMarkSizeLong = BenchMarkSize << 3
)

func generateIntSlice(b *testing.B) []int {
	var m = make([]int, BenchMarkSize)
	for i := range BenchMarkSize {
		m[i] = i
	}
	b.ResetTimer()
	return m
}

func generateUnsafePointerSlice(b *testing.B) [BenchMarkSize]unsafe.Pointer {
	var m [BenchMarkSize]unsafe.Pointer
	for i := range BenchMarkSize {
		e := i
		m[i] = unsafe.Pointer(&e)
	}
	b.ResetTimer()
	return m
}

func generateInterfaceSlice(b *testing.B) [BenchMarkSize]any {
	var m [BenchMarkSize]any
	for i := range BenchMarkSize {
		m[i] = i
	}
	b.ResetTimer()
	return m
}

func checkItem(b *testing.B, index, item int) {
	if item != index {
		b.Fail()
	}
}
