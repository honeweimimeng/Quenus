package batch

import (
	"github.com/honewemimeng/quenus/document"
	"github.com/honewemimeng/quenus/utils"
	"github.com/quentin-m/pqueue"
	"sync"
)

type PerDocBatch struct {
	factory DocBatchFactory
	pool    pqueue.PriorityQueue
	mu      sync.Mutex
}

func NewPerDocBatch(cap int, factory DocBatchFactory) *PerDocBatch {
	return &PerDocBatch{
		pool:    utils.NewFibonacciHeap(),
		factory: factory,
	}
}

func (b *PerDocBatch) ProcessDoc(doc *document.Document) {
	b.obtainLock()
	batch := b.obtainLock()
	batch.ProcessDoc(doc)
}

func (b *PerDocBatch) obtainLock() DocBatch {
	b.mu.Lock()
	defer b.mu.Unlock()
	inPool, _ := b.pool.Pop()
	if inPool == nil {
		return b.newDocBatch()
	}
	docBatch := inPool.(DocBatch)
	return docBatch
}

func (b *PerDocBatch) newDocBatch() DocBatch {
	docBatch, score := b.factory.NewDocBatch()
	b.pool.Push(docBatch, score)
	return docBatch
}

func (b *PerDocBatch) MaybeRelease(policy ReleasePolicy) bool {
	return false
}

func (b *PerDocBatch) Flush() {

}
