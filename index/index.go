package index

import "github.com/honewemimeng/quenus/directory"

type Index struct {
	meta           *Meta
	segmentManager *SegmentManager
	d              directory.Directory
}

func OpenIndex(d directory.Directory) *Index {
	i := &Index{d: d}
	i.Load()
	return i
}

func (i *Index) Load() {
	i.LoadMeta()
}

func (i *Index) LoadMeta() {
}

func (i *Index) Directory() directory.Directory {
	return i.d
}

func (i *Index) ReadSegment() {

}
