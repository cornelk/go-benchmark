# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.22.0 and 64 bit CPU on Linux using `make benchmark-perflock`.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
ValueUnsafePointer-8          23.52n ±  5%
ValueInterface-8              245.8n ± 13%
```

## Using defer vs not using it

```
Defer-8                       256.1n ±  7%
DeferNo-8                     124.9n ± 10%
```

## Iterating a slice

```
SliceReadRange-8              21.11n ±  5%
SliceReadForward-8            36.12n ±  6%
SliceReadBackwards-8          36.56n ±  6%
SliceReadLastItemFirst-8      36.38n ±  5%
```

## Passing a parameter by value vs pointer

```
ParameterPassedByPointer-8    10.94n ±  9%
ParameterPassedByValue-8      10.96n ±  4%
```

## Using reflect vs cast

```
Reflect-8                     344.8n ±  5%
Cast-8                        236.1n ±  7%
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
Hashing64MD5-8                140.5n ±  1%
Hashing64SHA1-8               170.4n ±  2%
Hashing64SHA256-8             124.9n ±  1%
Hashing64SHA3B224-8           697.2n ±  2%
Hashing64SHA3B256-8           693.5n ±  1%
Hashing64RIPEMD160-8          382.2n ±  2%
Hashing64Blake2B-8            377.4n ±  1%
Hashing64Blake2BSimd-8        347.4n ±  1%
Hashing64Murmur3-8            58.18n ±  2%
Hashing64Murmur3Twmb-8        57.77n ±  2%
Hashing64SipHash-8            59.51n ±  4%
Hashing64XXHash-8             36.08n ±  2%
Hashing64XXHashpier-8         41.67n ±  1%
Hashing64HighwayHash-8        96.17n ±  1%
```

## Filling a slice by index or append

```
SliceFillByIndex-8            15.88n ±  3%
SliceFillByIndexMake-8        15.83n ±  2%
SliceFillMakeAppend-8         40.36n ±  2%
SliceFillAppendNoMake-8       271.4n ±  6%
SliceFillSmallMakeAppend-8    338.2n ±  1%
```

## Writing and reading an int atomic

```
AtomicInt32-8                 956.8n ±  2%
AtomicInt64-8                 958.1n ±  3%
AtomicUintptr-8               951.5n ±  5%
```
