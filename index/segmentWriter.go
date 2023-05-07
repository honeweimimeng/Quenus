package index

import (
	"github.com/honewemimeng/quenus/analyzer"
	"github.com/honewemimeng/quenus/document"
	"github.com/honewemimeng/quenus/document/field"
)

type SegmentWriter struct {
	segment                *Segment
	perFieldPostingsWriter *PerFieldPostingsWriter
}

func NewSegmentWriter(segment *Segment) *SegmentWriter {
	return &SegmentWriter{
		segment: segment,
	}
}

func (w *SegmentWriter) WriterDoc(doc *document.Document) {
	schema := w.segment.idx.Schema()
	for _, indexField := range doc.Fields() {
		w.processField(indexField, schema, schema.GetAnalyzer(indexField))
	}
	w.segment.AfterUpdateDoc()
}

func (w *SegmentWriter) processField(field field.IndexField, schema *document.Schema, analyzer *analyzer.Analyzer) {
	proto := schema.GetFieldProto(field)
	if proto.IsText() {
		tokens := analyzer.Analyze(field.Val().(string))
		postingsWriter := w.perFieldPostingsWriter.FieldWriter(proto)
		postingsWriter.IndexText(w.segment.meta.MaxDoc, tokens)
	}
}
