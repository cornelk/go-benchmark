# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.19.0 and 64 bit CPU on Linux.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
BenchmarkValueUnsafePointer             38438962                31.28 ns/op
BenchmarkValueInterface                 35784254                34.47 ns/op
```

## Using defer vs not using it

```
BenchmarkDefer                           4565286               263.3 ns/op
BenchmarkDeferNo                        48542449                25.06 ns/op
```

## Iterating a slice

```
BenchmarkSliceReadRange                 66977254                18.43 ns/op
BenchmarkSliceReadForward               48649504                24.98 ns/op
BenchmarkSliceReadBackwards             28933062                44.26 ns/op
BenchmarkSliceReadLastItemFirst         42361663                31.40 ns/op
```

## Passing a parameter by value vs pointer

```
BenchmarkParameterPassedByPointer       167757066                7.093 ns/op
BenchmarkParameterPassedByValue         25323952                46.14 ns/op
```

## Using reflect vs cast

```
BenchmarkReflect                        10426641               114.5 ns/op
BenchmarkCast                           26002623                46.73 ns/op
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
BenchmarkHashing64MD5                   10458847               114.8 ns/op        69.71 MB/s
BenchmarkHashing64SHA1                   8493391               141.2 ns/op        56.66 MB/s
BenchmarkHashing64SHA256                 6398036               187.1 ns/op        42.75 MB/s
BenchmarkHashing64SHA3B224               2091546               573.1 ns/op        13.96 MB/s
BenchmarkHashing64SHA3B256               2102798               569.2 ns/op        14.06 MB/s
BenchmarkHashing64RIPEMD160              3957055               303.1 ns/op        26.40 MB/s
BenchmarkHashing64Blake2B                4042722               296.9 ns/op        26.95 MB/s
BenchmarkHashing64Blake2BSimd            4844781               247.7 ns/op        32.30 MB/s
BenchmarkHashing64Murmur3               25517697                47.19 ns/op      169.52 MB/s
BenchmarkHashing64Murmur3Twmb           25722096                46.34 ns/op      172.65 MB/s
BenchmarkHashing64SipHash               28485819                41.71 ns/op      191.82 MB/s
BenchmarkHashing64XXHash                46363768                25.71 ns/op      311.20 MB/s
BenchmarkHashing64XXHashpier            44888440                26.43 ns/op      302.64 MB/s
```

## Filling a slice by index or append

```
BenchmarkSliceFillByIndex               80740214                14.54 ns/op
BenchmarkSliceFillByIndexMake           80448338                14.51 ns/op
BenchmarkSliceFillMakeAppend            46544547                25.78 ns/op
BenchmarkSliceFillAppendNoMake           5809389               211.2 ns/op
BenchmarkSliceFillSmallMakeAppend        6118882               190.6 ns/op
```

## Writing and reading an int atomic

```
BenchmarkAtomicInt32                     1365267               878.1 ns/op
BenchmarkAtomicInt64                     1366141               878.4 ns/op
BenchmarkAtomicUintptr                   1367601               872.5 ns/op
```
