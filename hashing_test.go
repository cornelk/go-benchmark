package benchmark

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"testing"

	"github.com/cespare/xxhash"
	"github.com/dchest/siphash"
	blake2bsimd "github.com/minio/blake2b-simd"
	xxhash32pier "github.com/pierrec/xxHash/xxHash32"
	xxhash64pier "github.com/pierrec/xxHash/xxHash64"
	"github.com/spaolacci/murmur3"
	xxhash32vova "github.com/vova616/xxhash"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

const hashBufferSize = 8

func benchmarkHash(b *testing.B, hash func() hash.Hash, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)

	for i := 0; i < b.N; i++ {
		h := hash()
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}

func benchmarkHash64(b *testing.B, hash func() hash.Hash64, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	for i := 0; i < b.N; i++ {
		h := hash()
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}
func benchmarkHash64seed(b *testing.B, hash func(uint64) hash.Hash64, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	for i := 0; i < b.N; i++ {
		h := hash(1471)
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}
func benchmarkHash32seed(b *testing.B, hash func(uint32) hash.Hash32, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	for i := 0; i < b.N; i++ {
		h := hash(1471)
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}
func benchmarkHash64to32(b *testing.B, hash func() hash.Hash64, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	for i := 0; i < b.N; i++ {
		h := hash()
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		var b []byte
		s := h.Sum64()
		_ = append(
			b,
			byte(s>>56),
			byte(s>>48),
			byte(s>>40),
			byte(s>>32),
		)
	}
}
func benchmarkHash64to16(b *testing.B, hash func() hash.Hash64, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	for i := 0; i < b.N; i++ {
		h := hash()
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		var b []byte
		s := h.Sum64()
		_ = append(
			b,
			byte(s>>56),
			byte(s>>48),
		)
	}
}
func benchmarkHash64to8(b *testing.B, hash func() hash.Hash64, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	for i := 0; i < b.N; i++ {
		h := hash()
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		s := h.Sum64()
		_ = byte(s >> 56)
	}
}

func benchmarkHashKeyError(b *testing.B, hash func([]byte) (hash.Hash, error), length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	key := make([]byte, 16)

	for i := 0; i < b.N; i++ {
		h, _ := hash(key)
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}

func benchmarkHashKey64(b *testing.B, hash func([]byte) hash.Hash64, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	key := make([]byte, 16)

	for i := 0; i < b.N; i++ {
		h := hash(key)
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}
func benchmarkHashKey32(b *testing.B, hash func([]byte) hash.Hash32, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	key := make([]byte, 8)

	for i := 0; i < b.N; i++ {
		h := hash(key)
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}

func BenchmarkHashingMD5(b *testing.B) {
	benchmarkHash(b, md5.New, hashBufferSize)
}

func BenchmarkHashingSHA1(b *testing.B) {
	benchmarkHash(b, sha1.New, hashBufferSize)
}

func BenchmarkHashingSHA256(b *testing.B) {
	benchmarkHash(b, sha256.New, hashBufferSize)
}

func BenchmarkHashingSHA3B224(b *testing.B) {
	benchmarkHash(b, sha3.New224, hashBufferSize)
}

func BenchmarkHashingSHA3B256(b *testing.B) {
	benchmarkHash(b, sha3.New256, hashBufferSize)
}

func BenchmarkHashingRIPEMD160(b *testing.B) {
	benchmarkHash(b, ripemd160.New, hashBufferSize)
}

func BenchmarkHashingBlake2B(b *testing.B) {
	benchmarkHashKeyError(b, blake2b.New256, hashBufferSize)
}

func BenchmarkHashingBlake2BSimd(b *testing.B) {
	benchmarkHash(b, blake2bsimd.New256, hashBufferSize)
}

func BenchmarkHashingMurmur3(b *testing.B) {
	benchmarkHash64(b, murmur3.New64, hashBufferSize)
}
func BenchmarkHashingSipHash(b *testing.B) {
	benchmarkHashKey64(b, siphash.New, hashBufferSize)
}
func BenchmarkHashingXXHash64(b *testing.B) {
	benchmarkHash64(b, xxhash.New, hashBufferSize)
}
func BenchmarkHashingXXHash32vova(b *testing.B) {
	benchmarkHash32seed(b, xxhash32vova.New, hashBufferSize)
}
func BenchmarkHashingXXHash32pier(b *testing.B) {
	benchmarkHash32seed(b, xxhash32pier.New, hashBufferSize)
}
func BenchmarkHashingXXHash64pier(b *testing.B) {
	benchmarkHash64seed(b, xxhash64pier.New, hashBufferSize)
}
func BenchmarkHashingXXHash64to32(b *testing.B) {
	benchmarkHash64to32(b, xxhash.New, hashBufferSize)
}
func BenchmarkHashingXXHash64to16(b *testing.B) {
	benchmarkHash64to16(b, xxhash.New, hashBufferSize)
}
func BenchmarkHashingXXHash64to8(b *testing.B) {
	benchmarkHash64to8(b, xxhash.New, hashBufferSize)
}
