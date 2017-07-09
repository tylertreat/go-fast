package allocation

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

type Arena [][alloc]byte

func (a *Arena) Get() [alloc]byte {
	if len(*a) == 0 {
		*a = make([][alloc]byte, 1024)
	}

	b := &(*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return *b
}

func (a *Arena) Put(b [alloc]byte) {
	*a = append(*a, b)
}
