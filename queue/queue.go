package queue

import "sync/atomic"

// Queue concurrent queue
type Queue interface {
	Empty() bool
	Full() bool
	Pop() interface{}
	Push(interface{}) bool
	Size() (int, int)
}

type queueLockFree struct {
	capacity   int32
	data       []interface{}
	readIndex  int32
	writeIndex int32
}

// NewQueue to new a concurrent queue
func NewLockFreeQueue(cap int) Queue {
	return &queueLockFree{
		capacity:   int32(cap),
		data:       make([]interface{}, cap),
		readIndex:  0,
		writeIndex: 0,
	}
}

func (q *queueLockFree) Empty() bool {
	return q.readIndex == q.writeIndex
}

func (q *queueLockFree) Full() bool {
	return q.readIndex == nextIndex(q.writeIndex, q.capacity)
}

func (q *queueLockFree) Pop() interface{} {
	// double check for multi goroutine
	if q.Empty() {
		return nil
	}

	v := q.data[q.readIndex]
	q.readIndex = nextIndex(q.readIndex, q.capacity)
	return v
}

func (q *queueLockFree) Push(v interface{}) bool {
	var currentWriteIndex int32
	for {
		currentWriteIndex = q.writeIndex

		if q.readIndex == nextIndex(currentWriteIndex, q.capacity) {
			return false
		}

		if atomic.CompareAndSwapInt32(&q.writeIndex, currentWriteIndex, currentWriteIndex+1) {
			break
		}

	}

	q.data[currentWriteIndex] = v
	return true
}

func (q *queueLockFree) Size() (int, int) {
	return int(q.writeIndex - q.readIndex), cap(q.data)
}
