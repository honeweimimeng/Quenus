package index

import (
	"github.com/honeweimimeng/eventgo/driver"
	"github.com/honeweimimeng/eventgo/driver/event"
)

type Writer struct {
}

func NewWriter() *Writer {
	return &Writer{}
}

func (r *Writer) ReadDocHandle() event.SimpleHandler {
	return func(ch driver.Channel) {
		println(ch.Msg())
	}
}

func (r *Writer) EventHandle() ([]event.Proto, event.SimpleHandler) {
	return []event.Proto{NewWRITE(nil)}, r.ReadDocHandle()
}
