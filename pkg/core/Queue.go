package core

type QItem struct {
	value interface{}
	next  *QItem
}

type Queue struct {
	head *QItem
	tail *QItem
}

func (q *Queue) Push(value interface{}) {
	item := &QItem{value, nil}
	if q.tail != nil {
		q.tail.next = item
	}

	q.tail = item
	if q.head == nil {
		q.head = q.tail
	}
}

func (q *Queue) Pop() interface{} {
	if q.head == nil {
		panic("Queue is empty")
	}

	value := q.head.value
	q.head = q.head.next
	return value
}

func (q *Queue) CanPop() bool {
	return q.head != nil
}
