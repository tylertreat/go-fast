package channel

import (
	"sync"
	"testing"

	"github.com/Workiva/go-datastructures/queue"
)

func BenchmarkSimpleSet(b *testing.B) {
	set := make([]string, b.N)
	var mu sync.Mutex
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		set[i] = `a`
		mu.Unlock()
	}
}

func BenchmarkSimpleUnbufferedChannelSet(b *testing.B) {
	set := make([]string, b.N)
	c := make(chan string)
	go func() {
		i := 0
		for x := range c {
			set[i] = x
			i++
		}
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c <- `a`
	}
	close(c)
}

func BenchmarkSimpleBufferedChannelSet(b *testing.B) {
	set := make([]string, b.N)
	c := make(chan string, 1024)
	go func() {
		i := 0
		for x := range c {
			set[i] = x
			i++
		}
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c <- `a`
	}
	close(c)
}

func BenchmarkSimpleSetWriteContention(b *testing.B) {
	set := make([]string, 0, b.N)
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(8)
	b.ResetTimer()
	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				mu.Lock()
				set = append(set, `a`)
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkSimpleUnbufferedChannelSetWriteContention(b *testing.B) {
	set := make([]string, b.N)
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(8)
	go func() {
		i := 0
		for x := range c {
			set[i] = x
			i++
		}
	}()
	b.ResetTimer()
	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				c <- `a`
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c)
}

func BenchmarkSimpleBufferedChannelSetWriteContention(b *testing.B) {
	set := make([]string, b.N)
	c := make(chan string, 1024)
	var wg sync.WaitGroup
	wg.Add(8)
	go func() {
		i := 0
		for x := range c {
			set[i] = x
			i++
		}
	}()
	b.ResetTimer()
	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				c <- `a`
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c)
}
func BenchmarkBufferedChannel(b *testing.B) {
	ch := make(chan interface{}, 1024)
	var wg sync.WaitGroup
	wg.Add(1)

	b.ResetTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			<-ch
		}
		wg.Done()
	}()

	for i := 0; i < b.N; i++ {
		ch <- `a`
	}

	wg.Wait()
}

func BenchmarkRingBuffer(b *testing.B) {
	q := queue.NewRingBuffer(1024)
	var wg sync.WaitGroup
	wg.Add(1)

	b.ResetTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			q.Get()
		}
		wg.Done()
	}()

	for i := 0; i < b.N; i++ {
		q.Put(`a`)
	}

	wg.Wait()
}

func BenchmarkBufferedChannelReadContention(b *testing.B) {
	ch := make(chan interface{}, 1024)
	var wg sync.WaitGroup
	wg.Add(8)
	b.ResetTimer()

	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				<-ch
			}
			wg.Done()
		}()
	}

	for i := 0; i < b.N; i++ {
		ch <- `a`
	}

	wg.Wait()
}

func BenchmarkRingBufferReadContention(b *testing.B) {
	q := queue.NewRingBuffer(1024)
	var wg sync.WaitGroup
	wg.Add(8)
	b.ResetTimer()

	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				q.Get()
			}
			wg.Done()
		}()
	}

	for i := 0; i < b.N; i++ {
		q.Put(`a`)
	}

	wg.Wait()
}

func BenchmarkBufferedChannelWriteContention(b *testing.B) {
	ch := make(chan interface{}, 1024)
	var wg sync.WaitGroup
	wg.Add(8)
	if b.N < 8 {
		b.N = 8
	}
	b.ResetTimer()

	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				ch <- `a`
			}
			wg.Done()
		}()
	}

	for i := 0; i < b.N; i++ {
		<-ch
	}

	wg.Wait()
}

func BenchmarkRingBufferWriteContention(b *testing.B) {
	q := queue.NewRingBuffer(1024)
	var wg sync.WaitGroup
	wg.Add(8)
	if b.N < 8 {
		b.N = 8
	}
	b.ResetTimer()

	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				q.Put(`a`)
			}
			wg.Done()
		}()
	}

	for i := 0; i < b.N; i++ {
		q.Get()
	}

	wg.Wait()
}

func BenchmarkBufferedChannelReadWriteContention(b *testing.B) {
	ch := make(chan interface{}, 1024)
	var rwg sync.WaitGroup
	var wwg sync.WaitGroup
	rwg.Add(8)
	wwg.Add(8)
	b.ResetTimer()

	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				ch <- `a`
			}
			wwg.Done()
		}()
	}

	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				<-ch
			}
			rwg.Done()
		}()
	}

	wwg.Wait()
	rwg.Wait()
}

func BenchmarkRingBufferReadWriteContention(b *testing.B) {
	q := queue.NewRingBuffer(1024)
	var wwg sync.WaitGroup
	var rwg sync.WaitGroup
	wwg.Add(8)
	rwg.Add(8)
	b.ResetTimer()

	for i := 0; i < 8; i++ {
		go func() {
			for i := 0; i < b.N/8; i++ {
				q.Put(`a`)
			}
			wwg.Done()
		}()
	}

	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < b.N/8; j++ {
				q.Get()
			}
			rwg.Done()
		}()
	}

	wwg.Wait()
	rwg.Wait()
}
