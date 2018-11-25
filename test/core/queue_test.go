package core_test

import (
	"testing"

	. "github.com/FetchWeb/Fetch/pkg/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queue Test Suite")
}

func BenchmarkQueueSmallItem(b *testing.B) {
	var queue Queue

	// Benchmark pushing b.N items to queue, then separately popping
	// the same amount from the queue
	for n := 0; n < b.N; n++ {
		queueItem := QueueItem{Value: 42}
		queue.Push(queueItem)
	}

	for n := 0; n < b.N; n++ {
		queue.Pop()
	}
}

// The "LargeItem" is really not that large (about an extra 50 bytes or so), it really should
// be a bit bigger
func BenchmarkQueueLargeItem(b *testing.B) {
	var queue Queue

	// Benchmark pushing b.N items to queue, then separately popping
	// the same amount from the queue
	type largeItem struct {
		Number    int64
		FloatNum  float64
		Text      string
		Something interface{}
	}

	for n := 0; n < b.N; n++ {
		queueItem := QueueItem{Value: largeItem{Number: 123456789, FloatNum: 420.133769, Text: "This is a string for a large queue item", Something: 4.5}}
		queue.Push(queueItem)
	}

	for n := 0; n < b.N; n++ {
		queue.Pop()
	}
}

var _ = Describe("Queue", func() {
	var (
		queue      Queue
		queueItem1 QueueItem
		queueItem2 QueueItem
		queueItem3 QueueItem
		queueItem4 QueueItem
		queueItem5 QueueItem
	)

	BeforeEach(func() {
		queueItem1 = QueueItem{Value: "TEST"}
		queueItem2 = QueueItem{Value: 3.14}
		queueItem3 = QueueItem{Value: queueItem1}
		queueItem4 = QueueItem{Value: queue}
		queueItem5 = QueueItem{Value: "Something else to test for?"}
	})

	Context("Testing queue functionality", func() {
		Context("With one queue item", func() {
			It("Should only Pop once", func() {
				queue.Push(queueItem1)

				queue_count := 0
				for queue.CanPop() {
					queue.Pop()
					queue_count++
				}

				Expect(queue_count).To(Equal(1))
			})
		})

		Context("With five different queue items", func() {
			It("Should Pop all five items and they should be in FIFO order", func() {
				queue.Push(queueItem1)
				queue.Push(queueItem2)
				queue.Push(queueItem3)
				queue.Push(queueItem4)
				queue.Push(queueItem5)

				Expect(queue.CanPop()).To(Equal(true))
				Expect(queue.Pop()).To(Equal(queueItem1))

				Expect(queue.CanPop()).To(Equal(true))
				Expect(queue.Pop()).To(Equal(queueItem2))

				Expect(queue.CanPop()).To(Equal(true))
				Expect(queue.Pop()).To(Equal(queueItem3))

				Expect(queue.CanPop()).To(Equal(true))
				Expect(queue.Pop()).To(Equal(queueItem4))

				Expect(queue.CanPop()).To(Equal(true))
				Expect(queue.Pop()).To(Equal(queueItem5))

				Expect(queue.CanPop()).To(Equal(false))
			})
		})
	})
})
