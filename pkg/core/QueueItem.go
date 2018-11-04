package core

type QueueItem struct {
	Value interface{}
	next  *QueueItem
}
