package benchmark

import (
	"testing"
	"unsafe"
)

const (
	BenchMarkSize     = 1 << 10 // 1024
	BenchMarkSizeLong = BenchMarkSize << 5
)

func generateIntSlice(b *testing.B) [BenchMarkSize]int {
	var m [BenchMarkSize]int
	for i := 0; i < BenchMarkSize; i++ {
		m[i] = i
	}
	b.ResetTimer()
	return m
}

func generateUnsafePointerSlice(b *testing.B) [BenchMarkSize]unsafe.Pointer {
	var m [BenchMarkSize]unsafe.Pointer
	for i := 0; i < BenchMarkSize; i++ {
		e := i
		m[i] = unsafe.Pointer(&e)
	}
	b.ResetTimer()
	return m
}

func generateInterfaceSlice(b *testing.B) [BenchMarkSize]interface{} {
	var m [BenchMarkSize]interface{}
	for i := 0; i < BenchMarkSize; i++ {
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
