package index

import (
	"encoding/json"
	"github.com/honewemimeng/quenus/directory"
	"github.com/honewemimeng/quenus/document"
	"github.com/sirupsen/logrus"
)

const MetaFileName = "meta.json"

type Index struct {
	meta           *Meta
	dir            directory.Directory
	segmentManager *SegmentManager
}

func OpenIndex(d directory.Directory, schema *document.Schema, policy SegmentPolicy) *Index {
	i := &Index{dir: d, segmentManager: NewSegmentManager(policy)}
	exits, _ := d.Exists(MetaFileName, false)
	if exits {
		return i.LoadMeta()
	}
	return i.create(schema)
}

func (i *Index) Schema() *document.Schema {
	return i.meta.Schema
}

func (i *Index) Directory() directory.Directory {
	return i.dir
}

func (i *Index) create(schema *document.Schema) *Index {
	_, err := i.dir.Exists(MetaFileName, true)
	var segments []*SegmentMeta
	indexMeta := NewIndexMeta(segments, schema)
	err = indexMeta.SaveMeta(i.dir)
	if err != nil {
		panic(err.Error())
	}
	i.meta = indexMeta
	return i
}

func (i *Index) LoadMeta() *Index {
	fileOp := i.Directory().FileOp(MetaFileName, -1)
	metaData, err := i.Directory().Reader(fileOp).ReadAll()
	if err != nil || len(metaData) == 0 {
		panic(err.Error())
	}
	var indexMeta Meta
	if err := json.Unmarshal(metaData, &indexMeta); err != nil {
		panic(err.Error())
	}
	i.meta = &indexMeta
	return i
}

type IContext struct {
	logger *logrus.Logger
}

func NewIndexCtx() *IContext {
	return &IContext{logger: logrus.New()}
}
