package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	assert.Equal(t, 1, q.Pop())
}
