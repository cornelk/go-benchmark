# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.13.4 and 64 bit on Linux.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer-8           	23087859	        49.8 ns/op
BenchmarkValueInterface-8               	18875863	        61.7 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer-8                        	  480712	      2481 ns/op
BenchmarkDeferNo-8                      	12442630	        91 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange-8               	15874279	        73.1 ns/op
BenchmarkSliceReadForward-8             	12227348	        94.7 ns/op
BenchmarkSliceReadBackwards-8           	12207439	        94.4 ns/op
BenchmarkSliceReadLastItemFirst-8       	16149327	        71.8 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterPassedByPointer-8     	10073216	       121 ns/op
BenchmarkParameterPassedByValue-8       	 6407587	       193 ns/op
```

## Using reflect vs cast

```
BenchmarkReflect-8                      	 3802472	       275 ns/op
BenchmarkCast-8                         	12159952	        97 ns/op
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
BenchmarkHashing64MD5-8                 	 4124116	       254 ns/op	  31.46 MB/s
BenchmarkHashing64SHA1-8                	 3581850	       286 ns/op	  27.95 MB/s
BenchmarkHashing64SHA256-8              	 2776195	       381 ns/op	  20.97 MB/s
BenchmarkHashing64SHA3B224-8            	  828200	      1226 ns/op	   6.53 MB/s
BenchmarkHashing64SHA3B256-8            	 1079488	      1128 ns/op	   7.09 MB/s
BenchmarkHashing64RIPEMD160-8           	 1437418	       790 ns/op	  10.12 MB/s
BenchmarkHashing64Blake2B-8             	 2043148	       524 ns/op	  15.27 MB/s
BenchmarkHashing64Blake2BSimd-8         	 2176908	       498 ns/op	  16.07 MB/s
BenchmarkHashing64Murmur3-8             	11669984	        98.4 ns/op	  81.31 MB/s
BenchmarkHashing64Murmur3Twmb-8         	12181790	        94.7 ns/op	  84.46 MB/s
BenchmarkHashing64SipHash-8             	14001228	        83.6 ns/op	  95.70 MB/s
BenchmarkHashing64XXHash-8              	17448787	        59.5 ns/op	 134.56 MB/s
BenchmarkHashing64XXHashpier-8          	18139736	        72.4 ns/op	 110.45 MB/s
BenchmarkHashing64HighwayHash-8         	 7085976	       145 ns/op	  55.03 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex-8             	23782480	        48.1 ns/op
BenchmarkSliceFillByIndexMake-8         	46888129	        24.4 ns/op
BenchmarkSliceFillMakeAppend-8          	27952350	        40.5 ns/op
BenchmarkSliceFillAppendNoMake-8        	 1000000	      1112 ns/op
BenchmarkSliceFillSmallMakeAppend-8     	 1407986	       778 ns/op
```

## Writing and reading an int atomic

```
BenchmarkAtomicInt32-8                  	  266521	      4513 ns/op
BenchmarkAtomicInt64-8                  	  258920	      4510 ns/op
BenchmarkAtomicUintptr-8                	  257742	      4520 ns/op
```
