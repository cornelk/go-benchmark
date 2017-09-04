# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.9 and 64 bit on MacOS.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer-8              	 2000000	       667 ns/op
BenchmarkValueInterface-8                  	 1000000	      1061 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer-8                           	   20000	     65122 ns/op
BenchmarkDeferNo-8                         	  200000	      5587 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange-8                  	  300000	      3929 ns/op
BenchmarkSliceReadForward-8                	  500000	      3384 ns/op
BenchmarkSliceReadBackwards-8              	  500000	      3347 ns/op
BenchmarkSliceReadLastItemFirst-8          	  500000	      3319 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterSlicePassedByValue-8     	    5000	    277927 ns/op
BenchmarkParameterSlicePassedByPointer-8   	  500000	      3228 ns/op
```

## Using reflect vs cast

```
BenchmarkReflect-8                         	  300000	      4322 ns/op
BenchmarkCast-8                            	 1000000	      1674 ns/op
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
BenchmarkHashing64MD5-8           	       	5000000	       256 ns/op	  31.21 MB/s
BenchmarkHashing64SHA1-8           	       	5000000	       311 ns/op	  25.72 MB/s
BenchmarkHashing64SHA256-8           	    3000000	       569 ns/op	  14.04 MB/s
BenchmarkHashing64SHA3B224-8           	    1000000	      1190 ns/op	   6.72 MB/s
BenchmarkHashing64SHA3B256-8           	    1000000	      1195 ns/op	   6.69 MB/s
BenchmarkHashing64RIPEMD160-8           	    1000000	      1020 ns/op	   7.84 MB/s
BenchmarkHashing64Blake2B-8           	    2000000	       621 ns/op	  12.87 MB/s
BenchmarkHashing64Blake2BSimd-8           	3000000	       551 ns/op	  14.49 MB/s
BenchmarkHashing64Murmur3-8           	    20000000	       104 ns/op	  76.91 MB/s
BenchmarkHashing64SipHash-8           	    20000000	        92.9 ns/op	  86.12 MB/s
BenchmarkHashing64XXHash-8           	    20000000	        58.6 ns/op	 136.61 MB/s
BenchmarkHashing64XXHashpier-8           	20000000	        61.3 ns/op	 130.46 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex-8                	 2000000	       886 ns/op
BenchmarkSliceFillByIndexMake-8            	 2000000	       898 ns/op
BenchmarkSliceFillMakeAppend-8             	 2000000	       891 ns/op
```

## Writing and reading atomic ints

```
BenchmarkAtomicInt32-8     	                   20000	    363591 ns/op
BenchmarkAtomicInt64-8     	                   20000	    363817 ns/op
BenchmarkAtomicUintptr-8   	                   20000	    364249 ns/op
```
