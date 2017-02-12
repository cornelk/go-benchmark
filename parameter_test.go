package benchmark

import (
	"testing"
)

func getSliceItem(index int, m [128]int) int {
	return m[index]
}

func getSlicePointerItem(index int, m *[128]int) int {
	return m[index]
}

func BenchmarkParameterSlicePassedByValue(b *testing.B) {
	b.ReportAllocs()
	m := generateIntSlice()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 128; i++ {
				e := getSliceItem(i, m)
				if e != i {
					b.Fail()
				}
			}
		}
	})
}

func BenchmarkParameterSlicePassedByPointer(b *testing.B) {
	b.ReportAllocs()
	m := generateIntSlice()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 128; i++ {
				e := getSlicePointerItem(i, &m)
				if e != i {
					b.Fail()
				}
			}
		}
	})
}
