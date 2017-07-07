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

func BenchmarkSimpleChannelSet(b *testing.B) {
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
	b.ResetTimer()
	for i := 0; i < 4; i++ {
		for j := 0; j < b.N/4; j++ {
			mu.Lock()
			set = append(set, `a`)
			mu.Unlock()
		}
	}
}

func BenchmarkSimpleChannelSetWriteContention(b *testing.B) {
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
	for i := 0; i < 4; i++ {
		for j := 0; j < b.N/4; j++ {
			c <- `a`
		}
	}
	close(c)
}

func BenchmarkSimpleBufferedChannelSetWriteContention(b *testing.B) {
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
	for i := 0; i < 4; i++ {
		for j := 0; j < b.N/4; j++ {
			c <- `a`
		}
	}
	close(c)
}
func BenchmarkChannel(b *testing.B) {
	ch := make(chan interface{}, 1)

	b.ResetTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			<-ch
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- `a`
	}
}

func BenchmarkRingBuffer(b *testing.B) {
	q := queue.NewRingBuffer(1)

	b.ResetTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			q.Get()
		}
	}()

	for i := 0; i < b.N; i++ {
		q.Put(`a`)
	}
}

func BenchmarkLinkedList(b *testing.B) {
	q := NewLinkedList()

	b.ResetTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			q.Get()
		}
	}()

	for i := 0; i < b.N; i++ {
		q.Put(`a`)
	}
}

func BenchmarkChannelReadContention(b *testing.B) {
	ch := make(chan interface{}, 100)
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	go func() {
		for i := 0; i < b.N; i++ {
			ch <- `a`
		}
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N/1000; i++ {
				<-ch
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRingBufferReadContention(b *testing.B) {
	q := queue.NewRingBuffer(100)
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	go func() {
		for i := 0; i < b.N; i++ {
			q.Put(`a`)
		}
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N/1000; i++ {
				q.Get()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkLinkedListReadContention(b *testing.B) {
	q := NewLinkedList()
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	go func() {
		for i := 0; i < b.N; i++ {
			q.Put(`a`)
		}
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N/1000; i++ {
				q.Get()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkChannelContention(b *testing.B) {
	ch := make(chan interface{}, 100)
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				ch <- `a`
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				<-ch
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRingBufferContention(b *testing.B) {
	q := queue.NewRingBuffer(100)
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Put(`a`)
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Get()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkLinkedListContention(b *testing.B) {
	q := NewLinkedList()
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Put(`a`)
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Get()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
