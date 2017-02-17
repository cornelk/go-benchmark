package benchmark

import (
	"testing"
)

const BenchMarkSize = 1 << 10 // 1024

func generateIntSlice(b *testing.B) [BenchMarkSize]int {
	var m [BenchMarkSize]int
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
