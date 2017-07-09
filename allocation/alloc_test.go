package allocation

import (
	"sync"
	"testing"

	"github.com/couchbase/go-slab"
)

func BenchmarkSlab(b *testing.B) {
	arena := slab.NewArena(alloc>>1, alloc, 2, nil)
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = arena.Alloc(alloc)
		arena.DecRef(buf)
	}
	buf[0] = 0
}

func BenchmarkSlabParallel(b *testing.B) {
	arena := slab.NewArena(alloc>>1, alloc, 2, nil)
	var mu sync.Mutex
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var buf []byte
		for pb.Next() {
			mu.Lock()
			buf = arena.Alloc(alloc)
			arena.DecRef(buf)
			mu.Unlock()
		}
		if buf != nil {
			buf[0] = 0
		}
	})
}

func BenchmarkMake(b *testing.B) {
	b.ReportAllocs()
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf = make([]byte, alloc)
	}
	buf[0] = 0
}

func BenchmarkMakeParallel(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		var buf []byte
		for pb.Next() {
			buf = make([]byte, alloc)
		}
		if buf != nil {
			buf[0] = 0
		}
	})
}

var pool = sync.Pool{New: func() interface{} { return make([]byte, alloc) }}

func BenchmarkSyncPool(b *testing.B) {
	b.ReportAllocs()
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf = pool.Get().([]byte)
		pool.Put(buf)
	}
	buf[0] = 0
}

func BenchmarkSyncPoolParallel(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		var buf []byte
		for pb.Next() {
			buf = pool.Get().([]byte)
			pool.Put(buf)
		}
		if buf != nil {
			buf[0] = 0
		}
	})
}

func BenchmarkChannelPool(b *testing.B) {
	p := NewPool(1024)
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = p.Get()
		p.Put(buf)
	}
	buf[0] = 0
}

func BenchmarkChannelPoolParallel(b *testing.B) {
	p := NewPool(1024)
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf = p.Get()
			p.Put(buf)
		}
		if buf != nil {
			buf[0] = 0
		}
	})
}

func BenchmarkArena(b *testing.B) {
	arena := NewArena(1 << 30)
	var buf []byte
	b.N = 200000
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = arena.Alloc(alloc)
	}
	buf[0] = 0
}

func BenchmarkArenaParallel(b *testing.B) {
	arena := NewArena(1 << 30)
	b.N = 200000
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var buf []byte
		for pb.Next() {
			buf = arena.Alloc(alloc)
		}
		if buf != nil {
			buf[0] = 0
		}
	})
}
