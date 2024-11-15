# go-benchmark

Golang benchmarks used for optimizing code. The benchmarks were run with Golang 1.23.3 and 64 bit CPU on Linux using `make benchmark-perflock`.

## interface{} vs [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) 

```
ValueUnsafePointer-8          24.31n ±  3%
ValueInterface-8              216.2n ±  1%
```

## Using defer vs not using it

```
Defer-8                       237.5n ±  1%
DeferNo-8                     146.7n ±  1%
```

## Iterating a slice

```
SliceReadRange-8              24.21n ±  1%
SliceReadForward-8            40.07n ±  0%
SliceReadBackwards-8          40.54n ±  0%
SliceReadLastItemFirst-8      40.53n ±  7%
```

## Passing a parameter by value vs pointer

```
ParameterPassedByPointer-8    9.906n ±  0%
ParameterPassedByValue-8      9.950n ±  2%
```

## Using reflect vs cast

```
Reflect-8                     338.9n ±  0%
Cast-8                        216.4n ±  1%
```

## Hashing algorithms that produce a 64 bit hash of an 8 byte input

```
Hashing64MD5-8                156.4n ±  1%
Hashing64SHA1-8               186.3n ±  1%
Hashing64SHA256-8             134.5n ±  0%
Hashing64SHA3B224-8           470.1n ±  1%
Hashing64SHA3B256-8           469.9n ±  2%
Hashing64RIPEMD160-8          399.4n ±  1%
Hashing64Blake2B-8            420.8n ±  1%
Hashing64Blake2BSimd-8        360.5n ±  1%
Hashing64Murmur3-8            59.85n ±  1%
Hashing64Murmur3Twmb-8        60.65n ±  1%
Hashing64SipHash-8            63.14n ±  1%
Hashing64XXHash-8             37.37n ±  1%
Hashing64XXHashpier-8         42.93n ±  2%
Hashing64HighwayHash-8        94.79n ±  1%
```

## Filling a slice by index or append

```
SliceFillByIndex-8            19.20n ±  1%
SliceFillByIndexMake-8        19.21n ±  0%
SliceFillMakeAppend-8         30.91n ±  1%
SliceFillAppendNoMake-8       302.1n ± 10%
SliceFillSmallMakeAppend-8    336.9n ±  2%
```

## Writing and reading an int atomic

```
AtomicInt32-8                 1.144µ ±  1%
AtomicInt64-8                 1.145µ ±  1%
AtomicUintptr-8               1.145µ ±  0%
```
