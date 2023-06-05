package index

import (
	"github.com/honeweimimeng/eventgo/driver"
	"github.com/honeweimimeng/eventgo/driver/event"
	"github.com/honewemimeng/quenus/batch"
	"github.com/honewemimeng/quenus/document"
)

type Writer struct {
	idx              *Index
	docBatch         batch.DocBatch
	newSegmentPolicy batch.ReleasePolicy
	ctx              *IContext
}

func NewWriter(ctx *IContext, idx *Index) *Writer {
	perDocBatch := batch.NewPerDocBatch(ctx.perPoolCap, batch.NewZeroDocDataBatch())
	return &Writer{
		ctx:      ctx,
		idx:      idx,
		docBatch: perDocBatch,
	}
}

func (r *Writer) WriteDocHandle(doc *document.Document) {
	r.docBatch.ProcessDoc(doc)
	if r.docBatch.MaybeRelease(r.newSegmentPolicy) {
		r.idx.segmentManager.PubSegment()
	}
}

func (r *Writer) Commit() {
	r.docBatch.Flush()
	r.idx.segmentManager.PubSegment()
}

func (r *Writer) EventHandle() ([]event.Proto, event.SimpleHandler) {
	return []event.Proto{NewWRITE(nil)}, func(ch driver.Channel) {
		doc, ok := ch.Msg().(*document.Document)
		if !ok {
			r.ctx.logger.Println("cannot process channel in indexWriter,because channel msg is not type of Document")
		}
		r.WriteDocHandle(doc)
	}
}
