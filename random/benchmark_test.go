package random

import (
	crand "crypto/rand"
	mrand "math/rand"
	"testing"
)

func BenchmarkMathRand(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mrand.Float64()
	}
}
func BenchmarkCryptoRand(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := make([]byte, 8)
		_, _ = crand.Read(b)
	}
}
