# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.11.5 and 64 bit on Linux.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer-8           	30000000	        47.3 ns/op
BenchmarkValueInterface-8               	20000000	        65.1 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer-8                        	  500000	      2877 ns/op
BenchmarkDeferNo-8                      	 5000000	       270 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange-8               	10000000	       181 ns/op
BenchmarkSliceReadForward-8             	 5000000	       206 ns/op
BenchmarkSliceReadBackwards-8           	 5000000	       207 ns/op
BenchmarkSliceReadLastItemFirst-8       	10000000	       191 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterPassedByPointer-8     	10000000	       190 ns/op
BenchmarkParameterPassedByValue-8       	10000000	       194 ns/op
```

## Using reflect vs cast

```
BenchmarkReflect-8                      	 5000000	       258 ns/op
BenchmarkCast-8                         	20000000	        90.4 ns/op
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
BenchmarkHashing64MD5-8                 	 5000000	      2193 ns/op	   3.65 MB/s
BenchmarkHashing64SHA1-8                	 3000000	       406 ns/op	  19.69 MB/s
BenchmarkHashing64SHA256-8              	 3000000	       431 ns/op	  18.55 MB/s
BenchmarkHashing64SHA3B224-8            	 1000000	      1731 ns/op	   4.62 MB/s
BenchmarkHashing64SHA3B256-8            	 1000000	      1529 ns/op	   5.23 MB/s
BenchmarkHashing64RIPEMD160-8           	 1000000	      1049 ns/op	   7.63 MB/s
BenchmarkHashing64Blake2B-8             	 3000000	       609 ns/op	  13.12 MB/s
BenchmarkHashing64Blake2BSimd-8         	 2000000	       589 ns/op	  13.57 MB/s
BenchmarkHashing64Murmur3-8             	20000000	       114 ns/op	  70.14 MB/s
BenchmarkHashing64Murmur3Twmb-8         	10000000	       135 ns/op	  59.16 MB/s
BenchmarkHashing64SipHash-8             	20000000	        82.3 ns/op	  97.24 MB/s
BenchmarkHashing64XXHash-8              	20000000	        55.4 ns/op	 144.45 MB/s
BenchmarkHashing64XXHashpier-8          	20000000	        59.0 ns/op	 135.51 MB/s
BenchmarkHashing64HighwayHash-8         	10000000	       137 ns/op	  58.01 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex-8             	30000000	        45.6 ns/op
BenchmarkSliceFillByIndexMake-8         	50000000	        23.4 ns/op
BenchmarkSliceFillMakeAppend-8          	30000000	        42.0 ns/op
BenchmarkSliceFillAppendNoMake-8        	 3000000	      3907 ns/op
BenchmarkSliceFillSmallMakeAppend-8     	 2000000	       705 ns/op
```

## Writing and reading an int atomic

```
BenchmarkAtomicInt32-8                  	  300000	      4211 ns/op
BenchmarkAtomicInt64-8                  	  300000	      4226 ns/op
BenchmarkAtomicUintptr-8                	  300000	      4240 ns/op
```
