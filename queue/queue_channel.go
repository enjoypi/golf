package queue

type queueChannel struct {
	capacity int
	data     chan interface{}
}

func NewChannelQueue(cap int) Queue {
	return &queueChannel{
		capacity: cap,
		data:     make(chan interface{}, cap),
	}
}

func (q *queueChannel) Empty() bool {
	return len(q.data) <= 0
}

func (q *queueChannel) Full() bool {
	return len(q.data) >= cap(q.data)
}

func (q *queueChannel) Pop() interface{} {
	if q.Empty() {
		return nil
	}

	v := <-q.data
	return v
}

func (q *queueChannel) Push(v interface{}) bool {
	if q.Full() {
		return false
	}
	q.data <- v
	return true
}

func (q *queueChannel) Size() (int, int) {
	return len(q.data), q.capacity
}
