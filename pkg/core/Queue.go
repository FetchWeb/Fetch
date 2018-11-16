package core

import (
	"log"
)

type Queue struct {
	head *QueueItem
	tail *QueueItem
}

func (q *Queue) Push(value interface{}) {
	item := &QueueItem{value, nil}
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
		log.Fatal("Queue is empty, call CanPop() before calling Pop()")
		return nil
	}

	value := q.head.Value
	q.head = q.head.next

	if (q.head == nil) {
		q.tail = nil
	}
	return value
}

func (q *Queue) CanPop() bool {
	return q.head != nil
}
