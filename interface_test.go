package benchmark

import (
	"reflect"
	"testing"
)

// benchmark the use of unsafe.Pointer to store data
func BenchmarkValueUnsafePointer(b *testing.B) {
	m := generateUnsafePointerSlice(b)

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
	m := generateInterfaceSlice(b)

	for n := 0; n < b.N; n++ {
		for i := 0; i < BenchMarkSize; i++ {
			e := m[i]
			if e != i {
				b.Fail()
			}
		}
	}
}

func BenchmarkReflect(b *testing.B) {
	m := generateInterfaceSlice(b)

	for n := 0; n < b.N; n++ {
		for i := 0; i < BenchMarkSize; i++ {
			e := m[i]
			switch reflect.TypeOf(e).Kind() {
			case reflect.Int:
				if e != i {
					b.Fail()
				}
			default:
				b.Fail()
			}
		}
	}
}

func BenchmarkCast(b *testing.B) {
	m := generateInterfaceSlice(b)

	for n := 0; n < b.N; n++ {
		for i := 0; i < BenchMarkSize; i++ {
			e := m[i]
			switch e.(type) {
			case int:
				if e != i {
					b.Fail()
				}
			default:
				b.Fail()
			}
		}
	}
}
