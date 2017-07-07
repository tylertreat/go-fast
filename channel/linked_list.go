package channel

import "sync"

var nodePool = &sync.Pool{New: func() interface{} { return &node{} }}

type node struct {
	value interface{}
	next  *node
}

type LinkedList struct {
	head *node
	tail *node
	cond *sync.Cond
	mu   sync.Mutex
}

func NewLinkedList() *LinkedList {
	ll := &LinkedList{}
	ll.cond = sync.NewCond(&ll.mu)
	return ll
}

func (l *LinkedList) Put(v interface{}) {
	n := nodePool.Get().(*node)
	n.value = v
	n.next = nil
	l.mu.Lock()
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.cond.Signal()
	l.mu.Unlock()
}

func (l *LinkedList) Get() interface{} {
	l.mu.Lock()
	for l.head == nil {
		l.cond.Wait()
	}
	n := l.head
	l.head = n.next
	if l.head == nil {
		l.tail = nil
	}
	l.mu.Unlock()

	v := n.value
	nodePool.Put(n)
	return v
}
