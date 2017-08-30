# go-benchmark
Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.9 on MacOS.

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
BenchmarkHashingMD5-8                      	30000000	       247 ns/op	  32.29 MB/s
BenchmarkHashingSHA1-8           	       	20000000	       297 ns/op	  26.93 MB/s
BenchmarkHashingSHA256-8         	       	20000000	       546 ns/op	  14.65 MB/s
BenchmarkHashingSHA3B224-8              	    10000000	      1164 ns/op	   6.87 MB/s
BenchmarkHashingSHA3B256-8              	    10000000	      1177 ns/op	   6.79 MB/s
BenchmarkHashingRIPEMD160-8      	       	10000000	      1014 ns/op	   7.88 MB/s
BenchmarkHashingBlake2B-8        	       	10000000	       589 ns/op	  13.57 MB/s
BenchmarkHashingBlake2BSimd-8    	       	20000000	       524 ns/op	  15.24 MB/s
BenchmarkHashingMurmur3-8        	       	50000000	       133 ns/op	  60.06 MB/s
BenchmarkHashingSipHash-8        	       	50000000	       120 ns/op	  66.63 MB/s
BenchmarkHashingXXHash64-8       	       100000000	       89.5 ns/op	  89.36 MB/s
BenchmarkHashingXXHash64pier-8   	       100000000	       98.4 ns/op	  81.34 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex-8                	 2000000	       886 ns/op
BenchmarkSliceFillByIndexMake-8            	 2000000	       898 ns/op
BenchmarkSliceFillMakeAppend-8             	 2000000	       891 ns/op
```
