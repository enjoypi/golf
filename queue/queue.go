package queue

// Queue concurrent queue
type Queue interface {
	Empty() bool
	Full() bool
	Pop() interface{}
	Push(interface{}) bool
}

type queueLockFree struct {
	data       []interface{}
	readIndex  int
	writeIndex int
}

// NewQueue to new a concurrent queue
func NewLockFreeQueue(cap int) Queue {
	return &queueLockFree{data: make([]interface{}, 0, cap)}
}

func (q *queueLockFree) Empty() bool {
	return q.readIndex < q.writeIndex
}

func (q *queueLockFree) Full() bool {
	return false
}

func (q *queueLockFree) Pop() interface{} {
	if len(q.data) > 0 {
		v := q.data[0]
		q.data = q.data[1:]
		return v
	}
	return nil
}

func (q *queueLockFree) Push(v interface{}) bool {
	return false
}
