package benchmark

import (
	"testing"
)

func BenchmarkDefer(b *testing.B) {
	m := generateIntSlice(b)
	for n := 0; n < b.N; n++ {
		for i := 0; i < BenchMarkSize; i++ {
			func() {
				var item = m[i]
				defer checkItem(b, i, item)
			}()
		}
	}
}

func BenchmarkDeferNo(b *testing.B) {
	m := generateIntSlice(b)
	for n := 0; n < b.N; n++ {
		for i := 0; i < BenchMarkSize; i++ {
			func() {
				var item = m[i]
				checkItem(b, i, item)
			}()
		}
	}
}
