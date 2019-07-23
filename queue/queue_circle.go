package queue

import (
	"sync"
)
type queueCircleArray struct {
	capacity   int
	parallel   bool

	data       []interface{}
	mutex      sync.Mutex

	readIndex  int
	writeIndex int
}

func NewCircleArrayQueue(cap int, parallel bool) Queue {
	return &queueCircleArray{
		capacity:   cap,
		parallel : parallel,
		data:       make([]interface{}, cap),
		readIndex:  0,
		writeIndex: 0,
	}
}

func (q *queueCircleArray) Empty() bool {
	return q.readIndex == q.writeIndex
}

func (q *queueCircleArray) Full() bool {
	return q.readIndex == nextIndex(q.writeIndex, q.capacity)
}

func (q *queueCircleArray) Pop() interface{} {
	// double check for multi goroutine
	if q.Empty() {
		return nil
	}

	q.lock()

	if q.Empty() {
		q.unlock()
		return nil
	}

	v := q.data[q.readIndex]
	q.readIndex = nextIndex(q.readIndex, q.capacity)
	q.unlock()

	return v
}

func (q *queueCircleArray) Push(v interface{}) bool {
	// double check for multi goroutine
	if q.Full() {
		return false
	}

	q.lock()

	if q.Full() {
		q.unlock()
		return false
	}

	q.data[q.writeIndex] = v
	q.writeIndex = nextIndex(q.writeIndex, q.capacity)
	q.unlock()
	return true
}

func (q* queueCircleArray) lock()  {
	if q.parallel {
		q.mutex.Lock()
	}
}

func (q* queueCircleArray) unlock()  {
	if q.parallel {
		q.mutex.Unlock()
	}
}

func nextIndex(i int, cap int) int {
	if i+1 >= cap {
		return 0
	} else {
		return i + 1
	}
}
