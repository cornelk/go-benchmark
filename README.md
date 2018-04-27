# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.10.1 and 64 bit on MacOS.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer-8              	 2000000	       639 ns/op
BenchmarkValueInterface-8                  	 2000000	       962 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer-8                           	   20000	     58155 ns/op
BenchmarkDeferNo-8                         	  300000	      5378 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange-8                  	  500000	      3392 ns/op
BenchmarkSliceReadForward-8                	  500000	      3143 ns/op
BenchmarkSliceReadBackwards-8              	  500000	      3127 ns/op
BenchmarkSliceReadLastItemFirst-8          	  500000	      3140 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterSlicePassedByValue-8     	    5000	    268002 ns/op
BenchmarkParameterSlicePassedByPointer-8   	  500000	      3138 ns/op
```

## Using reflect vs cast

```
BenchmarkReflect-8                         	  500000	      4100 ns/op
BenchmarkCast-8                            	 1000000	      1562 ns/op
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
BenchmarkHashing64MD5-8                    	 5000000	       331 ns/op	  24.16 MB/s
BenchmarkHashing64SHA1-8                   	 5000000	       288 ns/op	  27.72 MB/s
BenchmarkHashing64SHA256-8                 	 3000000	       511 ns/op	  15.63 MB/s
BenchmarkHashing64SHA3B224-8               	 1000000	      1334 ns/op	   5.99 MB/s
BenchmarkHashing64SHA3B256-8               	 1000000	      1107 ns/op	   7.23 MB/s
BenchmarkHashing64RIPEMD160-8              	 2000000	       934 ns/op	   8.56 MB/s
BenchmarkHashing64Blake2B-8                	 3000000	       643 ns/op	  12.42 MB/s
BenchmarkHashing64Blake2BSimd-8            	 3000000	       568 ns/op	  14.08 MB/s
BenchmarkHashing64Murmur3-8                	20000000	       119 ns/op	  66.76 MB/s
BenchmarkHashing64SipHash-8                	20000000	        85.6 ns/op	  93.44 MB/s
BenchmarkHashing64XXHash-8                 	30000000	        64.7 ns/op	 123.74 MB/s
BenchmarkHashing64XXHashpier-8             	30000000	        69.4 ns/op	 115.25 MB/s
BenchmarkHashing64HighwayHash-8            	10000000	       147 ns/op	  54.27 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex-8                	 2000000	       870 ns/op
BenchmarkSliceFillByIndexMake-8            	 2000000	       879 ns/op
BenchmarkSliceFillMakeAppend-8             	 2000000	       880 ns/op
```

## Writing and reading an int atomic

```
BenchmarkAtomicInt32-8                     	    5000	    355971 ns/op
BenchmarkAtomicInt64-8                     	    5000	    354907 ns/op
BenchmarkAtomicUintptr-8                   	    3000	    353769 ns/op
```
