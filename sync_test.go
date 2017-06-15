package benchmark

import (
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

const (
	lokWrite = math.MinInt32 + 1
	lokRead  = int32(1)
)

type (
	SyncMutex struct {
		n  uint
		mu sync.RWMutex
	}
	SyncAtomic struct {
		n   uint
		lok int32
	}
)

func BenchmarkSyncRWMutex(b *testing.B) {
	m := &SyncMutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			m.mu.RLock()
			_ = m.n
			m.mu.RUnlock()

			m.mu.Lock()
			m.n++
			m.mu.Unlock()
		}
	})
}

func BenchmarkSyncRWAtomic(b *testing.B) {
	a := &SyncAtomic{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			atomic.AddInt32(&a.lok, lokRead)
			for atomic.LoadInt32(&a.lok) < lokRead {

			}
			_ = a.n
			atomic.AddInt32(&a.lok, -lokRead)

			for !atomic.CompareAndSwapInt32(&a.lok, 0, lokWrite) {

			}
			a.n++
			atomic.AddInt32(&a.lok, -lokWrite)
		}
	})
}

func BenchmarkSyncRWAtomicGosched(b *testing.B) {
	a := &SyncAtomic{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			atomic.AddInt32(&a.lok, lokRead)
			for atomic.LoadInt32(&a.lok) < lokRead {
				runtime.Gosched()
			}
			_ = a.n
			atomic.AddInt32(&a.lok, -lokRead)

			for !atomic.CompareAndSwapInt32(&a.lok, 0, lokWrite) {
				runtime.Gosched()
			}
			a.n++
			atomic.AddInt32(&a.lok, -lokWrite)
		}
	})
}
