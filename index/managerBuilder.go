package index

import (
	"github.com/honewemimeng/quenus/directory"
	"github.com/honewemimeng/quenus/document"
)

type ManagerBuilder struct {
	segmentPolicy SegmentPolicy
	schema        *document.Schema
}

func NewManagerBuilder() *ManagerBuilder {
	return &ManagerBuilder{}
}

func (b *ManagerBuilder) Policy(policy SegmentPolicy) {
	b.segmentPolicy = policy
}

func (b *ManagerBuilder) Schema(schema *document.Schema) {
	b.schema = schema
}

func (b *ManagerBuilder) Build(d directory.Directory) *Manager {
	return NewManager(b.LoadIndex(d), nil).Process()
}

func (b *ManagerBuilder) BuildAndProcess(d directory.Directory, next ManagerInitializer) *Manager {
	return NewManager(b.LoadIndex(d), NewInitializerChain(next)).Process()
}

func (b *ManagerBuilder) LoadIndex(d directory.Directory) *Index {
	policy := b.segmentPolicy
	if policy == nil {
		policy = NewDefaultSegmentPolicy()
	}
	return OpenIndex(d, b.schema, policy)
}
