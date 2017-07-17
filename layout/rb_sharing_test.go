package layout

import (
	"sync"
	"testing"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/stretchr/testify/assert"
)

func BenchmarkRBPaddedLifeCycle(b *testing.B) {
	rb := queue.NewRingBuffer(1024)

	var wwg sync.WaitGroup
	var rwg sync.WaitGroup
	wwg.Add(100)
	rwg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < b.N/100; j++ {
				_, err := rb.Get()
				assert.Nil(b, err)
			}
			rwg.Done()
		}()
	}

	b.ResetTimer()

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < b.N/100; j++ {
				rb.Put(i)
			}
			wwg.Done()
		}()
	}

	wwg.Wait()
	rwg.Wait()
}

func BenchmarkRBLifeCycle(b *testing.B) {
	rb := NewRingBuffer(1024) // Vendored without padding

	var wwg sync.WaitGroup
	var rwg sync.WaitGroup
	wwg.Add(100)
	rwg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < b.N/100; j++ {
				_, err := rb.Get()
				assert.Nil(b, err)
			}
			rwg.Done()
		}()
	}

	b.ResetTimer()

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < b.N/100; j++ {
				rb.Put(i)
			}
			wwg.Done()
		}()
	}

	wwg.Wait()
	rwg.Wait()
}
