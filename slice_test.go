package benchmark

import (
	"testing"
)

func BenchmarkSliceReadRange(b *testing.B) {
	m := generateIntSlice(b)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i, e := range m {
				checkItem(b, i, e)
			}
		}
	})
}

func BenchmarkSliceReadForward(b *testing.B) {
	m := generateIntSlice(b)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < BenchMarkSize; i++ {
				e := m[i]
				checkItem(b, i, e)
			}
		}
	})
}

func BenchmarkSliceReadBackwards(b *testing.B) {
	m := generateIntSlice(b)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := BenchMarkSize - 1; i >= 0; i-- {
				e := m[i]
				checkItem(b, i, e)
			}
		}
	})
}

func BenchmarkSliceReadLastItemFirst(b *testing.B) {
	m := generateIntSlice(b)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if m[BenchMarkSize-1] != BenchMarkSize-1 {
				b.Fail()
			}
			for i := 0; i < BenchMarkSize; i++ {
				e := m[i]
				checkItem(b, i, e)
			}
		}
	})
}
