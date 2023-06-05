package batch

import (
	"github.com/honewemimeng/quenus/document"
	"github.com/honewemimeng/quenus/document/field"
)

type ReleasePolicy interface {
}

type DocBatchFactory interface {
	NewDocBatch() (DocBatch, float64)
}

type DocBatch interface {
	ProcessDoc(doc *document.Document)
	MaybeRelease(policy ReleasePolicy) bool
	Flush()
}

type DocDataBatch struct {
	fieldWriter *PerFieldPostingWriter
}

func (b *DocDataBatch) ProcessDoc(doc *document.Document) {
	for i := 0; i < len(doc.Fields()); i++ {
		b.processField(doc.Fields()[i])
	}
}

func (b *DocDataBatch) processField(field field.IndexField) {
	println("===>")
}

func (b *DocDataBatch) MaybeRelease(policy ReleasePolicy) bool {
	return false
}

func (b *DocDataBatch) Flush() {

}

type ZeroDDBatchFactory struct {
}

func NewZeroDocDataBatch() *ZeroDDBatchFactory {
	return &ZeroDDBatchFactory{}
}

func (factory *ZeroDDBatchFactory) NewDocBatch() (DocBatch, float64) {
	return &DocDataBatch{fieldWriter: nil}, 0
}
