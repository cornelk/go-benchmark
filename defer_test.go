package benchmark

import (
	"testing"
)

// force both checks to not be inlined, otherwise the Go compiler would inline the checkItem call of the no-defer
// function, but currently the compiler can not inline functions that use defer.

//go:noinline
func checkDefer(b *testing.B, index, item int) {
	defer func() {
		if item != index {
			b.Fail()
		}
	}()
}

//go:noinline
func checkDeferNo(b *testing.B, index, item int) {
	if item != index {
		b.Fail()
	}
}

func BenchmarkDefer(b *testing.B) {
	m := generateIntSlice(b)
	for n := 0; n < b.N; n++ {
		for i := range m {
			var item = m[i]
			checkDefer(b, i, item)
		}
	}
}

func BenchmarkDeferNo(b *testing.B) {
	m := generateIntSlice(b)
	for n := 0; n < b.N; n++ {
		for i := range m {
			var item = m[i]
			checkDeferNo(b, i, item)
		}
	}
}
