# go-benchmark
Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.8 on MacOS.

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

## Hashing algorithms that produce a 64 bit hash of a short input

```
BenchmarkComparisonMD5-8                   	 5000000	       263 ns/op	  30.37 MB/s
BenchmarkComparisonSHA1-8                  	 5000000	       321 ns/op	  24.86 MB/s
BenchmarkComparisonSHA256-8                	 3000000	       569 ns/op	  14.04 MB/s
BenchmarkComparisonSHA3B224-8              	 1000000	      1245 ns/op	   6.42 MB/s
BenchmarkComparisonSHA3B256-8              	 1000000	      1262 ns/op	   6.33 MB/s
BenchmarkComparisonRIPEMD160-8             	 2000000	       997 ns/op	   8.02 MB/s
BenchmarkComparisonBlake2B-8               	 2000000	       632 ns/op	  12.64 MB/s
BenchmarkComparisonBlake2BSimd-8           	 3000000	       535 ns/op	  14.93 MB/s
BenchmarkComparisonMurmur3-8               	10000000	       139 ns/op	  57.23 MB/s
BenchmarkComparisonSipHash-8               	20000000	       126 ns/op	  63.40 MB/s
```
