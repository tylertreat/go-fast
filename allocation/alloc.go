package allocation

import (
	"sync/atomic"
	"unsafe"

	"github.com/Workiva/go-datastructures/queue"
)

const alloc = 1 << 12

// Pool holds byte arrays.
type Pool struct {
	pool chan []byte
}

// NewPool creates a new pool of byte array.
func NewPool(max int) *Pool {
	return &Pool{
		pool: make(chan []byte, max),
	}
}

// Get a byte array from the pool.
func (p *Pool) Get() []byte {
	select {
	case b := <-p.pool:
		return b
	default:
		return make([]byte, alloc)
	}
}

// Put returns a byte array to the pool.
func (p *Pool) Put(b []byte) {
	select {
	case p.pool <- b:
	default:
	}
}

// RBPool holds byte arrays.
type RBPool struct {
	pool *queue.RingBuffer
}

// NewRBPool creates a new pool of byte array.
func NewRBPool(max uint64) *RBPool {
	return &RBPool{
		pool: queue.NewRingBuffer(max),
	}
}

// Get a byte array from the pool.
func (p *RBPool) Get() []byte {
	if p.pool.Len() == 0 {
		return make([]byte, alloc)
	}
	b, _ := p.pool.Poll(1)
	if b == nil {
		b = make([]byte, alloc)
	}
	return b.([]byte)
}

// Put returns a byte array to the pool.
func (p *RBPool) Put(b []byte) {
	p.pool.Offer(b)
}

// Arena allacator.
type Arena struct {
	base uintptr
	data []byte
}

// NewArena creates a new arena allocator.
func NewArena(max int) *Arena {
	data := make([]byte, max, max)
	return &Arena{
		base: uintptr(unsafe.Pointer(&data[0])),
		data: data,
	}
}

// Alloc allocates a new byte slice from the arena.
func (arena *Arena) Alloc(size int) []byte {
	p := atomic.AddUintptr(&arena.base, uintptr(size))
	return (*[1 << 31]byte)(unsafe.Pointer(p))[:size]
}
