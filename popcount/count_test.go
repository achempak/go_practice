package popcount

import (
	"math/rand"
	"testing"
)

var x uint64 = rand.Uint64()

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(x)
	}
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountSlow(x)
	}
}

func BenchmarkPopCountClearRIght(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearRight(x)
	}
}
