package benchmark

import (
	"testing"
)

func BenchmarkSliceReadRange(b *testing.B) {
	m := generateIntSlice(b)
	for n := 0; n < b.N; n++ {
		for i, e := range m {
			checkItem(b, i, e)
		}
	}
}

func BenchmarkSliceReadForward(b *testing.B) {
	m := generateIntSlice(b)
	for n := 0; n < b.N; n++ {
		for i := 0; i < BenchMarkSize; i++ {
			e := m[i]
			checkItem(b, i, e)
		}
	}
}

func BenchmarkSliceReadBackwards(b *testing.B) {
	m := generateIntSlice(b)
	for n := 0; n < b.N; n++ {
		for i := BenchMarkSize - 1; i >= 0; i-- {
			e := m[i]
			checkItem(b, i, e)
		}
	}
}

func BenchmarkSliceReadLastItemFirst(b *testing.B) {
	m := generateIntSlice(b)
	for n := 0; n < b.N; n++ {
		if m[BenchMarkSize-1] != BenchMarkSize-1 {
			b.Fail()
		}
		for i := 0; i < BenchMarkSize; i++ {
			e := m[i]
			checkItem(b, i, e)
		}
	}
}

func BenchmarkSliceFillByIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var m [BenchMarkSize]int
		for i := 0; i < BenchMarkSize; i++ {
			m[i] = i
		}
	}
}

func BenchmarkSliceFillByIndexMake(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make([]int, BenchMarkSize)
		for i := 0; i < BenchMarkSize; i++ {
			m[i] = i
		}
	}
}

func BenchmarkSliceFillMakeAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make([]int, 0, BenchMarkSize)
		for i := 0; i < BenchMarkSize; i++ {
			m = append(m, i)
		}
	}
}
