# go-benchmark
Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.8.3 on MacOS.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer-8              	 2000000	       682 ns/op
BenchmarkValueInterface-8                  	 1000000	      1010 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer-8                           	   20000	     66176 ns/op
BenchmarkDeferNo-8                         	  200000	      5840 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange-8                  	  500000	      3779 ns/op
BenchmarkSliceReadForward-8                	  500000	      3286 ns/op
BenchmarkSliceReadBackwards-8              	  500000	      3267 ns/op
BenchmarkSliceReadLastItemFirst-8          	  500000	      3261 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterSlicePassedByValue-8     	    5000	    289891 ns/op
BenchmarkParameterSlicePassedByPointer-8   	  500000	      3343 ns/op
```

## Using reflect vs cast

```
BenchmarkReflect-8                         	  300000	      4279 ns/op
BenchmarkCast-8                            	 1000000	      1663 ns/op
```

## Hashing algorithms that produce a 64 bit hash of a short input

```
BenchmarkComparisonMD5-8                   	50000000	       262 ns/op	  30.42 MB/s
BenchmarkComparisonSHA1-8                  	50000000	       320 ns/op	  24.99 MB/s
BenchmarkComparisonSHA256-8                	30000000	       578 ns/op	  13.82 MB/s
BenchmarkComparisonSHA3B224-8              	10000000	      1226 ns/op	   6.52 MB/s
BenchmarkComparisonSHA3B256-8              	10000000	      1218 ns/op	   6.57 MB/s
BenchmarkComparisonRIPEMD160-8             	20000000	      1029 ns/op	   7.77 MB/s
BenchmarkComparisonBlake2B-8               	20000000	       615 ns/op	  12.99 MB/s
BenchmarkComparisonBlake2BSimd-8           	30000000	       539 ns/op	  14.82 MB/s
BenchmarkComparisonMurmur3-8               	100000000	       136 ns/op	  58.69 MB/s
BenchmarkComparisonSipHash-8               	100000000	       125 ns/op	  63.70 MB/s
BenchmarkComparisonXXHash-8                	200000000	        89 ns/op	  89.28 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex-4         	 3000000	       542 ns/op
BenchmarkSliceFillByIndexMake-4     	 2000000	       901 ns/op
BenchmarkSliceFillMakeAppend-4      	 1000000	      1309 ns/op
```
