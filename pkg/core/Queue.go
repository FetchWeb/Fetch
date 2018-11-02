package core

import (
	"log"
)

type QItem struct {
	Value interface{}
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
		log.Fatal("Queue is empty, call CanPop() before calling Pop()")
		return nil
	}

	value := q.head.Value
	q.head = q.head.next
	return value
}

func (q *Queue) CanPop() bool {
	return q.head != nil
}
