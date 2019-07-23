package queue

type queueCircleArray struct {
	data       []interface{}
	readIndex  int
	writeIndex int
	capacity   int
}

// NewQueue to new a concurrent queue
func NewCircleArrayQueue(cap int) Queue {
	return &queueCircleArray{
		data:       make([]interface{}, cap),
		readIndex:  0,
		writeIndex: 0,
		capacity:   cap,
	}
}

func (q *queueCircleArray) Empty() bool {
	return q.readIndex == q.writeIndex
}

func (q *queueCircleArray) Full() bool {
	return q.readIndex == nextIndex(q.writeIndex, q.capacity)
}

func (q *queueCircleArray) Pop() interface{} {
	if q.Empty() {
		return nil
	}

	v := q.data[q.readIndex]
	q.readIndex = nextIndex(q.readIndex, q.capacity)
	return v
}

func (q *queueCircleArray) Push(v interface{}) bool {
	if q.Full() {
		return false
	}

	q.data[q.writeIndex] = v
	q.writeIndex = nextIndex(q.writeIndex, q.capacity)
	return true
}

func nextIndex(i int, cap int) int {
	if i+1 >= cap {
		return 0
	} else {
		return i + 1
	}
}
