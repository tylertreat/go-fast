package allocation

import (
	"sync/atomic"
	"unsafe"
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

type Arena struct {
	base uintptr
	data []byte
}

func NewArena(max int) *Arena {
	data := make([]byte, max, max)
	return &Arena{
		base: uintptr(unsafe.Pointer(&data[0])),
		data: data,
	}
}

func (arena *Arena) Alloc(size int) []byte {
	p := atomic.AddUintptr(&arena.base, uintptr(size))
	return (*[1 << 31]byte)(unsafe.Pointer(p))[:size]
}
