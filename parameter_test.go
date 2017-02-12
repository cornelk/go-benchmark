package benchmark

import (
	"testing"
)

func getSliceItem(index int, m [BenchMarkSize]int) int {
	return m[index]
}

func getSlicePointerItem(index int, m *[BenchMarkSize]int) int {
	return m[index]
}

func BenchmarkParameterSlicePassedByValue(b *testing.B) {
	m := generateIntSlice(b)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < BenchMarkSize; i++ {
				e := getSliceItem(i, m)
				checkItem(b, i, e)
			}
		}
	})
}

func BenchmarkParameterSlicePassedByPointer(b *testing.B) {
	m := generateIntSlice(b)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < BenchMarkSize; i++ {
				e := getSlicePointerItem(i, &m)
				checkItem(b, i, e)
			}
		}
	})
}
