package benchmark

import (
	"testing"
)

func BenchmarkDefer(b *testing.B) {
	m := generateIntSlice(b)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < BenchMarkSize; i++ {
				func() {
					var item int
					item = m[i]
					defer checkItem(b, i, item)
				}()
			}
		}
	})
}

func BenchmarkDeferNo(b *testing.B) {
	m := generateIntSlice(b)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < BenchMarkSize; i++ {
				func() {
					var item int
					item = m[i]
					checkItem(b, i, item)
				}()
			}
		}
	})
}
