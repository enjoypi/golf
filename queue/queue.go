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
		data:       make([]interface{}, cap+1),
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
	for i := 1; i < 10; i++ {
		// to ensure thread-safety when there is more than 1 producer thread
		// a second index is defined (m_maximumReadIndex)
		currentReadIndex := q.readIndex

		if currentReadIndex == q.writeIndex {
			// the queue is empty or
			// a producer thread has allocate space in the queue but is
			// waiting to commit the data into it
			return nil
		}

		// retrieve the data from the queue
		v := q.data[currentReadIndex]

		// try to perfrom now the CAS operation on the read index. If we succeed
		// a_data already contains what m_readIndex pointed to before we
		// increased it
		if atomic.CompareAndSwapInt32(&q.readIndex, currentReadIndex, currentReadIndex+1) {
			return v
		}

		// it failed retrieving the element off the queue. Someone else must
		// have read the element stored at countToIndex(currentReadIndex)
		// before we could perform the CAS operation

	}

	return nil
}

func (q *queueLockFree) Push(v interface{}) bool {
	var currentWriteIndex int32
	for {
		currentWriteIndex = q.writeIndex

		// full
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
