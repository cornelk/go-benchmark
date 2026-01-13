# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.25.4 and 64 bit CPU on Linux using `make benchmark-perflock`.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
ValueUnsafePointer-8          24.52n ±  4%
ValueInterface-8              234.5n ±  1%
```

## Using defer vs not using it

```
Defer-8                       227.0n ±  1%
DeferNo-8                     145.8n ±  1%
```

## Iterating a slice

```
SliceReadRange-8              24.10n ±  2%
SliceReadForward-8            23.65n ±  1%
SliceReadBackwards-8          23.84n ±  2%
SliceReadLastItemFirst-8      24.29n ±  4%
```

## Passing a parameter by value vs pointer

```
ParameterPassedByPointer-8    12.05n ±  1%
ParameterPassedByValue-8      12.14n ±  1%
```

## Using reflect vs cast

```
Reflect-8                     340.6n ±  1%
Cast-8                        236.6n ±  4%
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
Hashing64MD5-8                151.5n ±  1%
Hashing64SHA1-8               127.4n ±  0%
Hashing64SHA256-8             129.7n ±  2%
Hashing64SHA3B224-8           488.8n ±  1%
Hashing64SHA3B256-8           485.7n ±  1%
Hashing64RIPEMD160-8          370.1n ±  3%
Hashing64Blake2B-8            415.5n ±  0%
Hashing64Blake2BSimd-8        352.6n ±  1%
Hashing64Murmur3-8            57.34n ±  1%
Hashing64Murmur3Twmb-8        57.16n ±  3%
Hashing64SipHash-8            62.06n ±  1%
Hashing64XXHash-8             34.33n ±  1%
Hashing64XXHashpier-8         39.35n ±  1%
Hashing64HighwayHash-8        88.59n ±  1%
```

## Filling a slice by index or append

```
SliceFillByIndex-8            18.88n ±  2%
SliceFillByIndexMake-8        18.89n ±  1%
SliceFillMakeAppend-8         43.14n ±  0%
SliceFillAppendNoMake-8       284.7n ± 15%
SliceFillSmallMakeAppend-8    355.0n ±  1%
```

## Writing and reading an int atomic

```
AtomicInt32-8                 1.144µ ±  0%
AtomicInt64-8                 1.135µ ±  0%
AtomicUintptr-8               1.132µ ±  1%
```
