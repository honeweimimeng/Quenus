package index

import (
	"github.com/honeweimimeng/eventgo/driver"
	"github.com/honeweimimeng/eventgo/driver/event"
	"github.com/honewemimeng/quenus/document"
)

type DocWR struct {
	DefaultIdxEvent
}

func NewWRITE(doc *document.Document) *DocWR {
	w := &DocWR{}
	w.name = "WRITE"
	w.Ch.EventProto = w
	w.Ch.MsgInstance = doc
	return w
}

func NewREAD(doc *document.Document) *DocWR {
	w := &DocWR{}
	w.name = "READ"
	w.Ch.EventProto = w
	w.Ch.MsgInstance = doc
	return w
}

type DefaultIdxEvent struct {
	name string
	id   int
	Ch   event.Channel
}

func (e *DefaultIdxEvent) Name() string {
	return e.name
}

func (e *DefaultIdxEvent) Id() int {
	return e.id
}

func (e *DefaultIdxEvent) Channel() driver.Channel {
	return &e.Ch
}
