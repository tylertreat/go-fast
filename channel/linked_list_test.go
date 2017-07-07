package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	assert := assert.New(t)
	ll := NewLinkedList()
	ll.Put(`a`)
	ll.Put(`b`)
	ll.Put(`c`)
	assert.Equal(`a`, ll.Get())
	assert.Equal(`b`, ll.Get())
	assert.Equal(`c`, ll.Get())
}
