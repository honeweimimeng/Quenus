package index

import (
	"github.com/honewemimeng/quenus/document"
)

type PostingsWriter interface {
	IndexText(docId DocId, terms []string)
	AddTermFreq(term string, docID DocId, termFreq uint64)
}

type PerFieldPostingsWriter struct {
	writer map[*document.FieldProto]PostingsWriter
}

func (p *PerFieldPostingsWriter) InitPerWriters(schema *document.Schema) {
	for _, proto := range schema.ProtoList() {
		p.writer[proto] = NewHashPostingWriter()
	}
}

func (p *PerFieldPostingsWriter) FieldWriter(proto *document.FieldProto) PostingsWriter {
	return p.writer[proto]
}

type HashPostingWriter struct {
	InvertedIndex   map[string][]DocId
	TermFrequencies map[string][]uint64
}

func NewHashPostingWriter() *HashPostingWriter {
	return &HashPostingWriter{
		InvertedIndex:   map[string][]DocId{},
		TermFrequencies: map[string][]uint64{},
	}
}

func (h *HashPostingWriter) IndexText(docId DocId, terms []string) {
	termFreqMap := map[string]uint64{}
	for _, term := range terms {
		termFreqMap[term] += 1
	}
	for term, freq := range termFreqMap {
		h.AddTermFreq(term, docId, freq)
	}
}

func (h *HashPostingWriter) AddTermFreq(term string, docID DocId, termFreq uint64) {
	h.InvertedIndex[term] = append(h.InvertedIndex[term], docID)
	h.TermFrequencies[term] = append(h.TermFrequencies[term], termFreq)
}
