package queue

import (
	"sync"
)

type queueWithLock struct {
	capacity   int
	data       []interface{}
	mutex      sync.Mutex
	readIndex  int
	writeIndex int
}

func NewQueueWithLock(cap int) Queue {
	return &queueWithLock{
		data:       make([]interface{}, cap),
		readIndex:  0,
		writeIndex: 0,
		capacity:   cap,
	}
}

func (q *queueWithLock) Empty() bool {
	return q.readIndex == q.writeIndex
}

func (q *queueWithLock) Full() bool {
	return q.readIndex == nextIndex(q.writeIndex, q.capacity)
}

func (q *queueWithLock) Pop() interface{} {
	if q.Empty() {
		return nil
	}

	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Empty() {
		return nil
	}

	readIndex := q.readIndex
	v := q.data[readIndex]
	q.readIndex = nextIndex(readIndex, q.capacity)
	return v
}

func (q *queueWithLock) Push(v interface{}) bool {
	if q.Full() {
		return false
	}

	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.Full() {
		return false
	}

	writeIndex := q.writeIndex
	q.data[writeIndex] = v

	q.writeIndex = nextIndex(writeIndex, q.capacity)
	return true
}
