package benchmark

import (
	"runtime"
	"sync"
	"testing"
)

// BenchmarkGoroutineNew
func BenchmarkGoroutineNew(b *testing.B) {
	vals := make([]int64, b.N)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg := &sync.WaitGroup{}
			wg.Add(b.N)
			cnt := int64(b.N)
			for n := int64(0); n < cnt; n++ {
				go calc(n, cnt, vals, wg)
			}
			wg.Wait()
		}
	})
}
func calc(n, cnt int64, vals []int64, wg *sync.WaitGroup) {
	vals[n] = n % cnt
	wg.Done()
}

// BenchmarkGoroutineChan
type (
	grp struct {
		vals  []int64
		valsC chan int64
		wg    *sync.WaitGroup
		cnt   int64
	}
)

func newGrp(N int64, valChanCnt int64, loopCnt int64) (r *grp) {
	r = &grp{
		vals:  make([]int64, N),
		valsC: make(chan int64, valChanCnt),
		wg:    &sync.WaitGroup{},
		cnt:   N,
	}
	for l := int64(0); l < loopCnt; l++ {
		go r.loop()
	}
	return r
}
func (g *grp) loop() {
	for n := range g.valsC {
		g.vals[n] = n % g.cnt
		g.wg.Done()
	}
}

func BenchmarkGoroutineChan1RteCPU(b *testing.B) {
	cnt := int64(b.N)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			grp := newGrp(cnt, 1, int64(runtime.NumCPU()))
			grp.wg.Add(b.N)
			for n := int64(0); n < cnt; n++ {
				grp.valsC <- n
			}
			grp.wg.Wait()
		}
	})
}
func BenchmarkGoroutineChan10RteCPU(b *testing.B) {
	cnt := int64(b.N)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			grp := newGrp(cnt, 10, int64(runtime.NumCPU()))
			grp.wg.Add(b.N)
			for n := int64(0); n < cnt; n++ {
				grp.valsC <- n
			}
			grp.wg.Wait()
		}
	})
}
func BenchmarkGoroutineChanCPURteCPU(b *testing.B) {
	cnt := int64(b.N)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			grp := newGrp(cnt, int64(runtime.NumCPU()), int64(runtime.NumCPU()))
			grp.wg.Add(b.N)
			for n := int64(0); n < cnt; n++ {
				grp.valsC <- n
			}
			grp.wg.Wait()
		}
	})
}
func BenchmarkGoroutineChan100RteCPU(b *testing.B) {
	cnt := int64(b.N)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			grp := newGrp(cnt, 100, int64(runtime.NumCPU()))
			grp.wg.Add(b.N)
			for n := int64(0); n < cnt; n++ {
				grp.valsC <- n
			}
			grp.wg.Wait()
		}
	})
}
func BenchmarkGoroutineChan10000RteCPU(b *testing.B) {
	cnt := int64(b.N)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			grp := newGrp(cnt, 10000, int64(runtime.NumCPU()))
			grp.wg.Add(b.N)
			for n := int64(0); n < cnt; n++ {
				grp.valsC <- n
			}
			grp.wg.Wait()
		}
	})
}
func BenchmarkGoroutineChan10Rte10(b *testing.B) {
	cnt := int64(b.N)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			grp := newGrp(cnt, 10, 10)
			grp.wg.Add(b.N)
			for n := int64(0); n < cnt; n++ {
				grp.valsC <- n
			}
			grp.wg.Wait()
		}
	})
}
func BenchmarkGoroutineChan100Rte100(b *testing.B) {
	cnt := int64(b.N)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			grp := newGrp(cnt, 100, 100)
			grp.wg.Add(b.N)
			for n := int64(0); n < cnt; n++ {
				grp.valsC <- n
			}
			grp.wg.Wait()
		}
	})
}
func BenchmarkGoroutineChan10000Rte10000(b *testing.B) {
	cnt := int64(b.N)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			grp := newGrp(cnt, 10000, 10000)
			grp.wg.Add(b.N)
			for n := int64(0); n < cnt; n++ {
				grp.valsC <- n
			}
			grp.wg.Wait()
		}
	})
}

// // BenchmarkGoroutineCond
// type (
// 	grpWithCond struct {
// 		vals []int64
// 		idx  uint32
// 		rtes []*rte
// 		wg   *sync.WaitGroup
// 		cnt  int64
// 	}
// 	rte struct {
// 		n    int64
// 		cond *sync.Cond
// 		g    *grpWithCond
// 	}
// )

// func newGrpWithCond(N int64, loopCnt int64) (r *grpWithCond) {
// 	r = &grpWithCond{
// 		vals: make([]int64, N),
// 		rtes: make([]*rte, int64(runtime.NumCPU())),
// 		wg:   &sync.WaitGroup{},
// 		cnt:  N,
// 	}
// 	for l := int64(0); l < loopCnt; l++ {
// 		r.rtes[l] = &rte{
// 			n:    math.MaxInt64,
// 			cond: &sync.Cond{L: &sync.Mutex{}},
// 			g:    r,
// 		}
// 		go r.rtes[l].loop()
// 	}
// 	return r
// }
// func (r *rte) loop() {
// 	for {
// 		r.cond.L.Lock()
// 		for r.n == math.MaxInt64 {
// 			r.cond.Wait()
// 		}
// 		//fmt.Printf("=-=-=- COMPUTING n:%v \n", r.n)
// 		r.g.vals[r.n] = r.n % r.g.cnt
// 		atomic.StoreInt64(&r.n, math.MaxInt64)
// 		r.g.wg.Done()
// 		r.cond.L.Unlock()
// 	}
// }

// func BenchmarkGoroutineCondGosched(b *testing.B) {
// 	cnt := int64(b.N)
// 	loopCnt := int64(runtime.NumCPU())
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			grp := newGrpWithCond(cnt, loopCnt)
// 			grp.wg.Add(b.N)
// 			for n := int64(0); n < cnt; n++ {
// 				runtime.Gosched()
// 				for !atomic.CompareAndSwapInt64(&grp.rtes[n%loopCnt].n, math.MaxInt64, n) {
// 					runtime.Gosched()
// 				}
// 				grp.rtes[n%loopCnt].cond.Signal()
// 			}
// 			grp.wg.Wait()
// 		}
// 	})
// }
