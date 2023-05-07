package index

import (
	"encoding/json"
	"github.com/honewemimeng/quenus/directory"
	"github.com/honewemimeng/quenus/document"
	uuid "github.com/honewemimeng/quenus/utils"
)

type DocId uint32

type SegmentId string

type Meta struct {
	Segments []*SegmentMeta   `json:"segments"`
	Schema   *document.Schema `json:"schema"`
}

func NewIndexMeta(segments []*SegmentMeta, schema *document.Schema) *Meta {
	return &Meta{
		Segments: segments,
		Schema:   schema,
	}
}

func (m *Meta) SaveMeta(d directory.Directory) error {
	metaJSON, err := json.Marshal(m)
	fileOp := d.FileOp(MetaFileName, -1)
	writer := d.Writer(fileOp)
	n, err := writer.WriteAt(metaJSON, 0)
	if err != nil || n == 0 {
		return &directory.IOErr{Err: "meta file write err"}
	}
	return nil
}

type SegmentMeta struct {
	Id     SegmentId `json:"id"`
	MaxDoc DocId     `json:"max_doc"`
}

func NewSegmentMeta(maxDocId DocId) *SegmentMeta {
	return &SegmentMeta{
		Id:     NewSegmentId(),
		MaxDoc: maxDocId,
	}
}

func NewSegmentId() SegmentId {
	return SegmentId(uuid.Generate())
}
