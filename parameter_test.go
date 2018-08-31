package main

import (
	"math/rand"
	"testing"
	"unsafe"
)

type Big struct {
	// Making the struct bigger will
	A [128]byte
}

func (b Big) byValue(x int) byte {
	r := byte(x)
	// Check every 32th element. This is faster
	// than checking each element but still triggers
	// reads on all cache lines as 32 is smaller than
	// the cache line size of modern CPUs.
	// We want this loop to be as fast as possible so the
	// percentual difference between our test cases is
	// as big as possible.
	for i := 0; i < len(b.A); i += 32 {
		r += b.A[i]
	}
	return r
}

func (b *Big) byPointer(x int) byte {
	r := byte(x)
	for i := 0; i < len(b.A); i += 32 {
		r += b.A[i]
	}
	return r
}

var (
	as []Big
)

func setup() {
	if as == nil {
		// Allocate 1GB of Big structs.
		as = make([]Big, 1024*1024*1024/unsafe.Sizeof(Big{}))

		// Make sure the memory isn't just zeroed pages and
		// is all mapped to our address space.
		// If we don't do this the first test case would be slower
		// as it is triggering all the page faults.
		for i := 0; i < len(as); i++ {
			for j := 0; j < len(as[i].A); j += 32 {
				as[i].A[j] = byte(rand.Intn(256))
			}
		}
	}
	rand.Seed(0)
}

func BenchmarkParameterPassedByPointer(b *testing.B) {
	setup()
	b.ResetTimer()

	// We randomly pick an element out of a 1GB big slice.
	// This way we assume the memory isn't already in the CPU cache.
	for i := 0; i < b.N; i++ {
		// Reading the value the pointer points to from memory will be slow on the first
		// byte but after that the memory will be in the CPU cache and following reads will be fast.
		as[rand.Intn(len(as))].byPointer(i)
	}
}

func BenchmarkParameterPassedByValue(b *testing.B) {
	setup()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Passing the parameter by value means Go will make a copy on the stack.
		// Reading the original value from memory will be slow but writing
		// the copy to the stack will be fast as it will be put into the CPU cache.
		// After this the function will read form the copy in the CPU cache which is fast again.
		// So while passing by pointer means extra instructions to dereference the pointer,
		// passing by value means extra writes to the CPU cache and extra CPU cache usage.
		// Since the CPU cache is very fast the difference between passing by pointer and
		// passing by value only gets significant with really big structs.
		//
		// In most cases passing by value should be preferred as it makes it easier for the Go escape
		// analysis to allocate the struct on the stack instead of the heap.
		// If passing by pointer makes your struct be allocated on the heap you'll get much
		// worse performance that passing by value.
		as[rand.Intn(len(as))].byValue(i)
	}
}
