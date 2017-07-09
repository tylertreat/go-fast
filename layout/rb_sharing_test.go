package layout

import (
	"sync"
	"testing"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/stretchr/testify/assert"
)

func BenchmarkRBPaddedLifeCycle(b *testing.B) {
	rb := queue.NewRingBuffer(64)

	var wwg sync.WaitGroup
	var rwg sync.WaitGroup
	wwg.Add(10)
	rwg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < b.N/10; j++ {
				_, err := rb.Get()
				assert.Nil(b, err)
			}
			rwg.Done()
		}()
	}

	b.ResetTimer()

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < b.N/10; j++ {
				rb.Put(i)
			}
			wwg.Done()
		}()
	}

	wwg.Wait()
	rwg.Wait()
}

func BenchmarkRBLifeCycle(b *testing.B) {
	rb := NewRingBuffer(64) // Vendored without padding

	var wwg sync.WaitGroup
	var rwg sync.WaitGroup
	wwg.Add(10)
	rwg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < b.N/10; j++ {
				_, err := rb.Get()
				assert.Nil(b, err)
			}
			rwg.Done()
		}()
	}

	b.ResetTimer()

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < b.N/10; j++ {
				rb.Put(i)
			}
			wwg.Done()
		}()
	}

	wwg.Wait()
	rwg.Wait()
}
