# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.20.6 and 64 bit CPU on Linux using `make benchmark-perflock`.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
ValueUnsafePointer-8          33.33n ±  4%
ValueInterface-8              180.2n ±  5%
```

## Using defer vs not using it

```
Defer-8                       284.2n ±  7%
DeferNo-8                     26.45n ±  1%
```

## Iterating a slice

```
SliceReadRange-8              19.71n ±  3%
SliceReadForward-8            25.89n ±  2%
SliceReadBackwards-8          46.48n ±  3%
SliceReadLastItemFirst-8      32.82n ±  4%
```

## Passing a parameter by value vs pointer

```
ParameterPassedByPointer-8    8.391n ±  5%
ParameterPassedByValue-8      8.755n ±  3%
```

## Using reflect vs cast

```
Reflect-8                     253.5n ±  5%
Cast-8                        181.1n ±  5%
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
Hashing64MD5-8                140.5n ±  1%
Hashing64SHA1-8               153.8n ±  2%
Hashing64SHA256-8             202.1n ±  2%
Hashing64SHA3B224-8           667.0n ±  1%
Hashing64SHA3B256-8           663.7n ±  1%
Hashing64RIPEMD160-8          328.6n ±  1%
Hashing64Blake2B-8            357.9n ±  3%
Hashing64Blake2BSimd-8        316.1n ±  3%
Hashing64Murmur3-8            55.27n ±  3%
Hashing64Murmur3Twmb-8        55.00n ±  1%
Hashing64SipHash-8            52.53n ±  1%
Hashing64XXHash-8             36.08n ±  1%
Hashing64XXHashpier-8         37.45n ±  1%
Hashing64HighwayHash-8        89.78n ±  1%
```

## Filling a slice by index or append

```
SliceFillByIndex-8            15.28n ±  0%
SliceFillByIndexMake-8        15.28n ±  0%
SliceFillMakeAppend-8         34.77n ±  1%
SliceFillAppendNoMake-8       253.5n ±  3%
SliceFillSmallMakeAppend-8    308.5n ±  1%
```

## Writing and reading an int atomic

```
AtomicInt32-8                 924.3n ±  1%
AtomicInt64-8                 923.9n ±  0%
AtomicUintptr-8               923.2n ±  0%
```
