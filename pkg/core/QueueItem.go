package core

// QueueItem stores an individual element of the linked list
type QueueItem struct {
	Value interface{}
	next  *QueueItem
}
