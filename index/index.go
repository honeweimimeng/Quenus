package index

import (
	"github.com/honeweimimeng/eventgo/driver/event"
	"github.com/honewemimeng/quenus/directory"
)

type Index struct {
	boot           event.BootStrap
	meta           *Meta
	segmentManager *SegmentManager
	d              directory.Directory
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
