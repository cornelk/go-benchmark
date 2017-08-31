package benchmark

import (
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicInt32(b *testing.B) {
	var a int32
	for n := 0; n < b.N; n++ {
		a = 0
		for i := int32(0); i < BenchMarkSizeLong; i++ {
			atomic.AddInt32(&a, 1)
			if atomic.LoadInt32(&a) != (i + 1) {
				b.Fail()
			}
		}
	}
}

func BenchmarkAtomicInt64(b *testing.B) {
	var a int64
	for n := 0; n < b.N; n++ {
		a = 0
		for i := int64(0); i < BenchMarkSizeLong; i++ {
			atomic.AddInt64(&a, 1)
			if atomic.LoadInt64(&a) != (i + 1) {
				b.Fail()
			}
		}
	}
}

func BenchmarkAtomicUintptr(b *testing.B) {
	var a uintptr
	for n := 0; n < b.N; n++ {
		a = 0
		for i := uintptr(0); i < BenchMarkSizeLong; i++ {
			atomic.AddUintptr(&a, 1)
			if atomic.LoadUintptr(&a) != (i + 1) {
				b.Fail()
			}
		}
	}
}
