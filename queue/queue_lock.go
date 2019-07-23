package queue

import (
	"sync"
)

type queueWithLock struct {
	data []interface{}
	sync.Mutex
}

// NewQueue to new a concurrent queue
func NewQueueWithLock() Queue {
	return &queueWithLock{data: make([]interface{}, 0, 256)}
}

func (q *queueWithLock) Empty() bool {
	return len(q.data) <= 0
}

func (q *queueWithLock) Pop() interface{} {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if len(q.data) > 0 {
		v := q.data[0]
		q.data = q.data[1:]
		return v
	}
	return nil
}

func (q *queueWithLock) Push(v interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.data = append(q.data, v)
}
