package index

import (
	"github.com/honeweimimeng/eventgo/driver/event"
	"github.com/honewemimeng/quenus/directory"
	"github.com/honewemimeng/quenus/document"
	"github.com/sirupsen/logrus"
)

type Index struct {
	schema *document.Schema
	boot   event.BootStrap
	meta   *Meta
	d      directory.Directory
	ctx    *IContext
}

func OpenIndex(d directory.Directory) *Index {
	i := &Index{d: d}
	return i
}

func (i *Index) Directory() directory.Directory {
	return i.d
}

func (i *Index) ReadSegment() {

}

func (i *Index) LoadMeta() *Index {
	return i
}

func (i *Index) LoadSegment() *Index {
	return i
}

type IContext struct {
	logger *logrus.Logger
}

func NewIndexCtx() *IContext {
	return &IContext{logger: logrus.New()}
}
