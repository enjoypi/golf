package queue

type Queue interface {
	Push(interface{})
	Pop() interface{}
}

type qk struct {
}

func NewQueue() Queue {
	return &qk{}
}

func (q *qk) Push(interface{}) {

}

func (q *qk) Pop() interface{} {
	return nil
}
