package benchmark

import (
	"testing"
	"unsafe"
)

// benchmark the use of unsafe.Pointer to store data
func BenchmarkValueUnsafePointer(b *testing.B) {
	b.ReportAllocs()
	var m [128]unsafe.Pointer
	for i := 0; i < 128; i++ {
		e := i
		m[i] = unsafe.Pointer(&e)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 128; i++ {
				item := m[i]
				e := *(*int)(item)
				if e != i {
					b.Fail()
				}
			}
		}
	})
}

// benchmark the use of interface{} to store data
func BenchmarkValueInterface(b *testing.B) {
	b.ReportAllocs()
	var m [128]interface{}
	for i := 0; i < 128; i++ {
		var e interface{}
		e = i
		m[i] = e
	}
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
