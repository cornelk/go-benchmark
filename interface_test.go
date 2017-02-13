package benchmark

import (
	"testing"
	"unsafe"
)

// benchmark the use of unsafe.Pointer to store data
func BenchmarkValueUnsafePointer(b *testing.B) {
	var m [BenchMarkSize]unsafe.Pointer
	for i := 0; i < BenchMarkSize; i++ {
		e := i
		m[i] = unsafe.Pointer(&e)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < BenchMarkSize; i++ {
			item := m[i]
			e := *(*int)(item)
			if e != i {
				b.Fail()
			}
		}
	}
}

// benchmark the use of interface{} to store data
func BenchmarkValueInterface(b *testing.B) {
	var m [BenchMarkSize]interface{}
	for i := 0; i < BenchMarkSize; i++ {
		var e interface{}
		e = i
		m[i] = e
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < BenchMarkSize; i++ {
			e := m[i]
			if e != i {
				b.Fail()
			}
		}
	}
}
