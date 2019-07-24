package queue

// Queue concurrent queue
type Queue interface {
	Empty() bool
	Full() bool
	Pop() interface{}
	Push(interface{}) bool
}

type queueLockFree struct {
	capacity   int
	data       []interface{}
	readIndex  int
	writeIndex int
}

// NewQueue to new a concurrent queue
func NewLockFreeQueue(cap int) Queue {
	return &queueLockFree{
		capacity:   cap,
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
	// double check for multi goroutine
	if q.Full() {
		return false
	}

	q.data[q.writeIndex] = v
	q.writeIndex = nextIndex(q.writeIndex, q.capacity)
	return true
}
