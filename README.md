# go-benchmark
Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.7.5 on MacOS.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer-8              	 2000000	       704 ns/op
BenchmarkValueInterface-8                  	  100000	     12030 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer-8                           	   10000	    110894 ns/op
BenchmarkDeferNo-8                         	  200000	      6596 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange-8                  	  300000	      4312 ns/op
BenchmarkSliceReadForward-8                	  500000	      3769 ns/op
BenchmarkSliceReadBackwards-8              	  500000	      3708 ns/op
BenchmarkSliceReadLastItemFirst-8          	  500000	      3779 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterSlicePassedByValue-8     	    5000	    300823 ns/op
BenchmarkParameterSlicePassedByPointer-8   	  300000	      3773 ns/op
```

## Hashing algorithms that produce a 64 bit hash of a short input

```
BenchmarkComparisonMD5-8                   	 5000000	       254 ns/op	  31.41 MB/s
BenchmarkComparisonSHA1-8                  	 5000000	       309 ns/op	  25.81 MB/s
BenchmarkComparisonSHA256-8                	 3000000	       575 ns/op	  13.89 MB/s
BenchmarkComparisonSHA3B224-8              	 1000000	      1264 ns/op	   6.33 MB/s
BenchmarkComparisonSHA3B256-8              	 1000000	      1235 ns/op	   6.48 MB/s
BenchmarkComparisonRIPEMD160-8             	 1000000	      1066 ns/op	   7.50 MB/s
BenchmarkComparisonBlake2B-8               	 2000000	       618 ns/op	  12.94 MB/s
BenchmarkComparisonBlake2BSimd-8           	 3000000	       546 ns/op	  14.65 MB/s
BenchmarkComparisonMurmur3-8               	10000000	       138 ns/op	  57.59 MB/s
BenchmarkComparisonSipHash-8               	10000000	       133 ns/op	  59.79 MB/s
```
