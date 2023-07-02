package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

var (
	counter int64
	mu      sync.Mutex
)

func BenchmarkMutex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

var (
	counter2 int64
)

func BenchmarkAtomic(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&counter2, 1)
	}
}
