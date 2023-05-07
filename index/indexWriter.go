package index

import (
	"github.com/honeweimimeng/eventgo/driver"
	"github.com/honeweimimeng/eventgo/driver/event"
	"github.com/honewemimeng/quenus/document"
)

type Writer struct {
	idx *Index
	ctx *IContext
}

func NewWriter(ctx *IContext, idx *Index) *Writer {
	return &Writer{ctx: ctx, idx: idx}
}

func (r *Writer) WriteDocHandle(doc *document.Document) {
	segment := r.idx.segmentManager.GetSegment(r.idx)
	writer := NewSegmentWriter(segment)
	writer.WriterDoc(doc)
}

func (r *Writer) Commit() {
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
