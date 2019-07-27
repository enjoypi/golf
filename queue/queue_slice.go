package queue

type sliceQueue struct {
	capacity int
	data     []interface{}
}

func NewSliceQueue(cap int) Queue {
	return &sliceQueue{
		capacity: cap,
		data:     make([]interface{}, 0, cap),
	}
}

func (q *sliceQueue) Empty() bool {
	return len(q.data) <= 0
}

func (q *sliceQueue) Full() bool {
	return len(q.data) >= q.capacity
}

func (q *sliceQueue) Pop() interface{} {
	if q.Empty() {
		return nil
	}

	v := q.data[0]
	q.data = q.data[1:]

	return v
}

func (q *sliceQueue) Push(v interface{}) bool {
	if q.Full() {
		return false
	}

	q.data = append(q.data, v)
	return true
}

func (q *sliceQueue) Size() (int, int) {
	return len(q.data), cap(q.data)
}
