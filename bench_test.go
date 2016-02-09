package farm

import "testing"

var res32 uint32
var res64 uint64
var res64lo, res64hi uint64

var buf = make([]byte, 256)

func BenchmarkHash32(b *testing.B) {
	var r uint32
	for i := 0; i < b.N; i++ {
		// record the result to prevent the compiler eliminating the function call
		r = Hash32(buf)
	}
	// store the result to a package level variable so the compiler cannot eliminate the Benchmark itself
	res32 = r
}

func BenchmarkHash64(b *testing.B) {
	var r uint64
	for i := 0; i < b.N; i++ {
		r = Hash64(buf)
	}
	res64 = r
}

func BenchmarkHash128(b *testing.B) {
	var rlo, rhi uint64
	for i := 0; i < b.N; i++ {
		rlo, rhi = Hash128(buf)
	}
	res64lo = rlo
	res64hi = rhi
}
