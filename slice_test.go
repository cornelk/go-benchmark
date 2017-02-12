package benchmark

import (
	"testing"
)

func BenchmarkSliceReadRange(b *testing.B) {
	m := generateIntSlice()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i, e := range m {
				if e != i {
					b.Fail()
				}
			}
		}
	})
}

func BenchmarkSliceReadForward(b *testing.B) {
	m := generateIntSlice()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 128; i++ {
				e := m[i]
				if e != i {
					b.Fail()
				}
			}
		}
	})
}

func BenchmarkSliceReadBackwards(b *testing.B) {
	m := generateIntSlice()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 127; i >= 0; i-- {
				e := m[i]
				if e != i {
					b.Fail()
				}
			}
		}
	})
}

func BenchmarkSliceReadLastItemFirst(b *testing.B) {
	m := generateIntSlice()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if m[127] != 127 {
				b.Fail()
			}
			for i := 0; i < 128; i++ {
				e := m[i]
				if e != i {
					b.Fail()
				}
			}
		}
	})
}
