package index

import (
	"github.com/honeweimimeng/eventgo/driver"
	"github.com/honeweimimeng/eventgo/driver/event"
)

type Reader struct {
	idx *Index
	ctx *IContext
}

func NewReader(ctx *IContext, idx *Index) *Reader {
	return &Reader{ctx: ctx, idx: idx}
}

func (r *Reader) WriteDocHandle() event.SimpleHandler {
	return func(ch driver.Channel) {
		println("===>read", ch.Msg())
	}
}

func (r *Reader) EventHandle() ([]event.Proto, event.SimpleHandler) {
	return []event.Proto{NewREAD(nil)}, r.WriteDocHandle()
}
