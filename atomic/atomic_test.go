package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkCheckMutex(b *testing.B) {
	var (
		mu    sync.Mutex
		check bool
	)
	for i := 0; i < b.N; i++ {
		mu.Lock()
		c := check
		mu.Unlock()
		if c {
		}
	}
}

func BenchmarkCheckAtomic(b *testing.B) {
	var check int32
	for i := 0; i < b.N; i++ {
		if atomic.LoadInt32(&check) == 1 {
		}
	}
}

func BenchmarkCheckMutexContention(b *testing.B) {
	var (
		mu    sync.Mutex
		check bool
		wg    sync.WaitGroup
	)
	wg.Add(1)
	b.ResetTimer()

	go func() {
		for i := 0; i < b.N; i++ {
			mu.Lock()
			check = i%10 == 0
			mu.Unlock()
		}
		wg.Done()
	}()

	for i := 0; i < b.N; i++ {
		mu.Lock()
		c := check
		mu.Unlock()
		if c {
		}
	}

	wg.Wait()
}

func BenchmarkCheckAtomicContention(b *testing.B) {
	var (
		check int32
		wg    sync.WaitGroup
	)
	wg.Add(1)

	go func() {
		for i := 0; i < b.N; i++ {
			atomic.StoreInt32(&check, int32(i%10))
		}
		wg.Done()
	}()

	for i := 0; i < b.N; i++ {
		if atomic.LoadInt32(&check) == 1 {
		}
	}

	wg.Wait()
}
