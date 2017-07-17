package layout

import (
	"sync"
	"testing"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/stretchr/testify/assert"
)

func BenchmarkRBPaddedLifeCycle(b *testing.B) {
	rb := queue.NewRingBuffer(1024)

	var wg sync.WaitGroup
	wg.Add(100)
	b.ResetTimer()

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < b.N; j++ {
				_, err := rb.Get()
				assert.Nil(b, err)
			}
			wg.Done()
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < b.N; j++ {
				rb.Put(i)
			}
		}()
	}

	wg.Wait()
}

func BenchmarkRBLifeCycle(b *testing.B) {
	rb := NewRingBuffer(1024) // Vendored without padding

	var wg sync.WaitGroup
	wg.Add(100)
	b.ResetTimer()

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < b.N; j++ {
				_, err := rb.Get()
				assert.Nil(b, err)
			}
			wg.Done()
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < b.N; j++ {
				rb.Put(i)
			}
		}()
	}

	wg.Wait()
}
