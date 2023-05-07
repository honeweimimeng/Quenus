package document

import (
	"github.com/honewemimeng/quenus/analyzer"
	"github.com/honewemimeng/quenus/document/field"
)

type FieldProto struct {
	text         bool
	field        field.IndexField
	analyzerName string
}

func (p *FieldProto) IsText() bool {
	return p.text
}

func (p *FieldProto) IsBin() bool {
	return !p.text
}

type Schema struct {
	fieldSchemaMap map[int]*FieldProto
	analyzerMap    map[string]*analyzer.Analyzer
	fieldProto     []*FieldProto
}

func NewSchema() *Schema {
	return &Schema{}
}

func (s *Schema) ProtoList() []*FieldProto {
	return s.fieldProto
}

func (s *Schema) GetFieldProto(fi field.IndexField) *FieldProto {
	return s.fieldSchemaMap[fi.Id()]
}

func (s *Schema) GetAnalyzer(fi field.IndexField) *analyzer.Analyzer {
	proto := s.GetFieldProto(fi)
	analyze := s.analyzerMap[proto.analyzerName]
	return analyze
}
