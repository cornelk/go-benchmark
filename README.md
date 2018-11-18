# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.11.2 and 64 bit on MacOS.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer-8             	 2000000	       664 ns/op
BenchmarkValueInterface-8                 	 2000000	       997 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer-8                          	   20000	     63183 ns/op
BenchmarkDeferNo-8                        	  200000	      6026 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange-8                 	  500000	      3264 ns/op
BenchmarkSliceReadForward-8               	  500000	      3585 ns/op
BenchmarkSliceReadBackwards-8             	  500000	      3267 ns/op
BenchmarkSliceReadLastItemFirst-8         	  500000	      3286 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterPassedByPointer-8       	10000000	       157 ns/op
BenchmarkParameterPassedByValue-8         	10000000	       159 ns/op
```

## Using reflect vs cast

```
BenchmarkReflect-8                        	  300000	      4383 ns/op
BenchmarkCast-8                           	 1000000	      1638 ns/op
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
BenchmarkHashing64MD5-8                   	 5000000	       330 ns/op	  24.19 MB/s
BenchmarkHashing64SHA1-8                  	 5000000	       401 ns/op	  19.94 MB/s
BenchmarkHashing64SHA256-8                	 3000000	       529 ns/op	  15.11 MB/s
BenchmarkHashing64SHA3B224-8              	 1000000	      1406 ns/op	   5.69 MB/s
BenchmarkHashing64SHA3B256-8              	 1000000	      1169 ns/op	   6.84 MB/s
BenchmarkHashing64RIPEMD160-8             	 2000000	       689 ns/op	  11.60 MB/s
BenchmarkHashing64Blake2B-8               	 3000000	       743 ns/op	  10.76 MB/s
BenchmarkHashing64Blake2BSimd-8           	 3000000	       794 ns/op	  10.07 MB/s
BenchmarkHashing64Murmur3-8               	10000000	       107 ns/op	  74.76 MB/s
BenchmarkHashing64SipHash-8               	20000000	       118 ns/op	  67.73 MB/s
BenchmarkHashing64XXHash-8                	20000000	        59.6 ns/op	 134.19 MB/s
BenchmarkHashing64XXHashpier-8            	20000000	        95.2 ns/op	  84.05 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex-8               	 5000000	       358 ns/op
BenchmarkSliceFillByIndexMake-8           	 5000000	       356 ns/op
BenchmarkSliceFillMakeAppend-8            	 2000000	       926 ns/op
```

## Writing and reading an int atomic

```
BenchmarkAtomicInt32-8                    	    5000	    365611 ns/op
BenchmarkAtomicInt64-8                    	    5000	    359317 ns/op
BenchmarkAtomicUintptr-8                  	    5000	    363894 ns/op
```
