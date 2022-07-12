# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.18.3 and 64 bit CPU on Linux.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer       	63332010	        17.58 ns/op
BenchmarkValueInterface           	33523322	        33.75 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer                    	 4621748	       260.7 ns/op
BenchmarkDeferNo                  	36943111	        41.77 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange           	39174472	        30.51 ns/op
BenchmarkSliceReadForward         	39017761	        43.46 ns/op
BenchmarkSliceReadBackwards       	46495941	        30.26 ns/op
BenchmarkSliceReadLastItemFirst   	38621402	        31.26 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterPassedByPointer 	186839847	         6.393 ns/op
BenchmarkParameterPassedByValue   	187472469	         6.397 ns/op
```

## Using reflect vs cast

```
BenchmarkReflect                  	 9228655	       126.9 ns/op
BenchmarkCast                     	26482546	        45.60 ns/op
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
BenchmarkHashing64MD5             	10712892	       111.7 ns/op	  71.63 MB/s
BenchmarkHashing64SHA1            	 9161918	       130.9 ns/op	  61.11 MB/s
BenchmarkHashing64SHA256          	 6586072	       182.3 ns/op	  43.89 MB/s
BenchmarkHashing64SHA3B224        	 2254980	       532.2 ns/op	  15.03 MB/s
BenchmarkHashing64SHA3B256        	 2264166	       525.5 ns/op	  15.22 MB/s
BenchmarkHashing64RIPEMD160       	 3958806	       303.2 ns/op	  26.39 MB/s
BenchmarkHashing64Blake2B         	 4217047	       286.0 ns/op	  27.97 MB/s
BenchmarkHashing64Blake2BSimd     	 4676388	       256.0 ns/op	  31.25 MB/s
BenchmarkHashing64Murmur3         	29881615	        39.89 ns/op	 200.57 MB/s
BenchmarkHashing64Murmur3Twmb     	30463790	        39.29 ns/op	 203.61 MB/s
BenchmarkHashing64SipHash         	30221272	        39.09 ns/op	 204.64 MB/s
BenchmarkHashing64XXHash          	49449735	        24.05 ns/op	 332.62 MB/s
BenchmarkHashing64XXHashpier      	48240354	        24.49 ns/op	 326.61 MB/s
BenchmarkHashing64HighwayHash     	21095451	        56.72 ns/op	 141.05 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex         	81669337	        14.48 ns/op
BenchmarkSliceFillByIndexMake     	82515597	        14.55 ns/op
BenchmarkSliceFillMakeAppend      	37082814	        32.52 ns/op
BenchmarkSliceFillAppendNoMake    	 6654816	       192.4 ns/op
BenchmarkSliceFillSmallMakeAppend 	 5722377	       188.4 ns/op
```

## Writing and reading an int atomic

```
BenchmarkAtomicInt32              	 1388169	       869.3 ns/op
BenchmarkAtomicInt64              	 1372506	       870.2 ns/op
BenchmarkAtomicUintptr            	 1370652	       868.9 ns/op
```
