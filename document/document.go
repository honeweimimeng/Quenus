package document

import (
	"github.com/honewemimeng/quenus/document/field"
	"strings"
)

type Document struct {
	Cap    int
	fields []field.IndexField
}

func NewDocument(cap int) *Document {
	return &Document{Cap: cap}
}

func (d *Document) AddField(field field.IndexField) {
	d.fields = append(d.fields, field)
}

func (d *Document) Fields() []field.IndexField {
	return d.fields
}

func (d *Document) GetField(name string) field.IndexField {
	var r field.IndexField
	for _, item := range d.fields {
		if strings.EqualFold(name, item.Name()) {
			return item
		}
	}
	return r
}

func (d *Document) GetFieldByIndex(i int) field.IndexField {
	return d.fields[i]
}
