package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	var (
		counter int64
		mu      sync.Mutex
	)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			go func() {
				mu.Lock()
				counter++
				mu.Unlock()
			}()
		}
	})
}

func BenchmarkAtomic(b *testing.B) {
	var counter int64

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			go func() {
				atomic.AddInt64(&counter, 1)
			}()
		}
	})
}
