package index

import (
	"github.com/honeweimimeng/eventgo/driver"
	"github.com/honeweimimeng/eventgo/driver/event"
)

type Reader struct {
}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) WriteDocHandle() event.SimpleHandler {
	return func(ch driver.Channel) {

	}
}

func (r *Reader) EventHandle() ([]event.Proto, event.SimpleHandler) {
	return []event.Proto{NewREAD(nil)}, r.WriteDocHandle()
}
