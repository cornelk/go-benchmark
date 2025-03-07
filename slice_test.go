package benchmark

import (
	"container/list"
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
		for i := range BenchMarkSize {
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
		for i := range BenchMarkSize {
			e := m[i]
			checkItem(b, i, e)
		}
	}
}

func BenchmarkSliceFillByIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var m [BenchMarkSize]int
		for i := range BenchMarkSize {
			m[i] = i
		}
	}
}

func BenchmarkSliceFillByIndexMake(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make([]int, BenchMarkSize)
		for i := range BenchMarkSize {
			m[i] = i
		}
	}
}

func BenchmarkSliceFillMakeAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make([]int, 0, BenchMarkSize)
		for i := range BenchMarkSize {
			m = append(m, i)
		}
	}
}

func BenchmarkSliceFillAppendNoMake(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var m []int

		for i := range BenchMarkSize {
			m = append(m, i)
		}
	}
}

func BenchmarkSliceFillSmallMakeAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make([]int, 0, BenchMarkSize/10)
		for i := range BenchMarkSize {
			m = append(m, i)
		}
	}
}

func BenchmarkFillLinkedListPushBack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := list.New()
		for range BenchMarkSize {
			l.PushBack("a")
		}
	}
}

func BenchmarkFillLinkedListPushFront(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := list.New()
		for range BenchMarkSize {
			l.PushFront("a")
		}
	}
}
