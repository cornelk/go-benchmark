# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.9.2 and 64 bit on MacOS.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer-8              	 2000000	       805 ns/op
BenchmarkValueInterface-8                  	 1000000	      1391 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer-8                           	   20000	     63285 ns/op
BenchmarkDeferNo-8                         	  300000	      5525 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange-8                  	  300000	      4516 ns/op
BenchmarkSliceReadForward-8                	  300000	      4115 ns/op
BenchmarkSliceReadBackwards-8              	  300000	      4234 ns/op
BenchmarkSliceReadLastItemFirst-8          	  300000	      4206 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterSlicePassedByValue-8     	    3000	    396628 ns/op
BenchmarkParameterSlicePassedByPointer-8   	  300000	      4172 ns/op
```

## Using reflect vs cast

```
BenchmarkReflect-8                         	  200000	      5966 ns/op
BenchmarkCast-8                            	  500000	      2381 ns/op
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
BenchmarkHashing64MD5-8                    	 5000000	       363 ns/op	  22.00 MB/s
BenchmarkHashing64SHA1-8                   	 3000000	       379 ns/op	  21.07 MB/s
BenchmarkHashing64SHA256-8                 	 2000000	       647 ns/op	  12.36 MB/s
BenchmarkHashing64SHA3B224-8               	 1000000	      1937 ns/op	   4.13 MB/s
BenchmarkHashing64SHA3B256-8               	 1000000	      1572 ns/op	   5.09 MB/s
BenchmarkHashing64RIPEMD160-8              	 1000000	      1312 ns/op	   6.09 MB/s
BenchmarkHashing64Blake2B-8                	 2000000	       693 ns/op	  11.54 MB/s
BenchmarkHashing64Blake2BSimd-8            	 2000000	       628 ns/op	  12.73 MB/s
BenchmarkHashing64Murmur3-8                	10000000	       127 ns/op	  62.57 MB/s
BenchmarkHashing64SipHash-8                	10000000	       114 ns/op	  70.05 MB/s
BenchmarkHashing64XXHash-8                 	20000000	        93 ns/op	  86.12 MB/s
BenchmarkHashing64XXHashpier-8             	20000000	       105 ns/op	  75.62 MB/s
BenchmarkHashing64HighwayHash-8            	10000000	       282 ns/op	  28.31 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex-8                	 1000000	      1208 ns/op
BenchmarkSliceFillByIndexMake-8            	 1000000	      1222 ns/op
BenchmarkSliceFillMakeAppend-8             	 1000000	      1244 ns/op
```

## Writing and reading an int atomic

```
BenchmarkAtomicInt32-8                     	    5000	    351130 ns/op
BenchmarkAtomicInt64-8                     	    5000	    355952 ns/op
BenchmarkAtomicUintptr-8                   	    5000	    348614 ns/op
```
