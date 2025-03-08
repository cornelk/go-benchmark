# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.24.0 and 64 bit CPU on Linux using `make benchmark-perflock`.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
ValueUnsafePointer-8          24.81n ±  3%
ValueInterface-8              233.9n ±  0%
```

## Using defer vs not using it

```
Defer-8                       228.3n ±  1%
DeferNo-8                     145.6n ±  0%
```

## Iterating a slice

```
SliceReadRange-8              23.79n ±  3%
SliceReadForward-8            23.82n ±  2%
SliceReadBackwards-8          23.70n ±  1%
SliceReadLastItemFirst-8      24.61n ±  3%
```

## Passing a parameter by value vs pointer

```
ParameterPassedByPointer-8    12.01n ±  0%
ParameterPassedByValue-8      12.04n ±  0%
```

## Using reflect vs cast

```
Reflect-8                     340.3n ±  0%
Cast-8                        240.9n ±  3%
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
Hashing64MD5-8                152.9n ±  0%
Hashing64SHA1-8               185.7n ±  0%
Hashing64SHA256-8             132.6n ±  1%
Hashing64SHA3B224-8           474.2n ±  1%
Hashing64SHA3B256-8           472.7n ±  1%
Hashing64RIPEMD160-8          389.0n ±  1%
Hashing64Blake2B-8            418.0n ±  1%
Hashing64Blake2BSimd-8        362.7n ±  1%
Hashing64Murmur3-8            58.68n ±  1%
Hashing64Murmur3Twmb-8        56.82n ±  1%
Hashing64SipHash-8            60.16n ±  1%
Hashing64XXHash-8             34.35n ±  2%
Hashing64XXHashpier-8         39.04n ±  1%
Hashing64HighwayHash-8        94.09n ±  1%
```

## Filling a slice by index or append

```
SliceFillByIndex-8            18.90n ±  0%
SliceFillByIndexMake-8        18.94n ±  0%
SliceFillMakeAppend-8         42.85n ±  0%
SliceFillAppendNoMake-8       278.6n ±  6%
SliceFillSmallMakeAppend-8    360.1n ±  7%
```

## Writing and reading an int atomic

```
AtomicInt32-8                 1.150µ ±  0%
AtomicInt64-8                 1.140µ ±  0%
AtomicUintptr-8               1.138µ ±  0%
```
