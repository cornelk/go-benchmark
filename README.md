# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.21.0 and 64 bit CPU on Linux using `make benchmark-perflock`.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
ValueUnsafePointer-8           20.62n ± 3%
ValueInterface-8               189.0n ± 3%
```

## Using defer vs not using it

```
Defer-8                        220.2n ± 3%
DeferNo-8                      119.5n ± 3%
```

## Iterating a slice

```
SliceReadRange-8               19.83n ± 2%
SliceReadForward-8             33.35n ± 2%
SliceReadBackwards-8           33.12n ± 8%
SliceReadLastItemFirst-8       33.20n ± 3%
```

## Passing a parameter by value vs pointer

```
ParameterPassedByPointer-8     9.396n ± 4%
ParameterPassedByValue-8       9.737n ± 6%
```

## Using reflect vs cast

```
Reflect-8                      287.2n ± 6%
Cast-8                         182.8n ± 5%
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
Hashing64MD5-8                 143.6n ± 5%
Hashing64SHA1-8                155.3n ± 6%
Hashing64SHA256-8              116.8n ± 1%
Hashing64SHA3B224-8            662.2n ± 8%
Hashing64SHA3B256-8            664.1n ± 2%
Hashing64RIPEMD160-8           328.9n ± 2%
Hashing64Blake2B-8             369.8n ± 2%
Hashing64Blake2BSimd-8         326.3n ± 1%
Hashing64Murmur3-8             55.04n ± 1%
Hashing64Murmur3Twmb-8         54.62n ± 3%
Hashing64SipHash-8             54.35n ± 1%
Hashing64XXHash-8              35.32n ± 1%
Hashing64XXHashpier-8          37.06n ± 1%
Hashing64HighwayHash-8         83.62n ± 1%
```

## Filling a slice by index or append

```
SliceFillByIndex-8             15.51n ± 0%
SliceFillByIndexMake-8         15.50n ± 0%
SliceFillMakeAppend-8          25.52n ± 4%
SliceFillAppendNoMake-8        253.1n ± 5%
SliceFillSmallMakeAppend-8     311.0n ± 4%
```

## Writing and reading an int atomic

```
AtomicInt32-8                  917.3n ± 1%
AtomicInt64-8                  922.9n ± 1%
AtomicUintptr-8                924.4n ± 1%
```
