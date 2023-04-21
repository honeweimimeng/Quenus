package index

import (
	"github.com/honewemimeng/quenus/directory"
)

type ManagerBuilder struct {
	directory directory.Directory
}

func NewManagerBuilder(d directory.Directory) *ManagerBuilder {
	return &ManagerBuilder{directory: d}
}

func (b *ManagerBuilder) Build() *Manager {
	return NewManager(nil).Process()
}

func (b *ManagerBuilder) BuildAndProcess(next ManagerInitializer) *Manager {
	return NewManager(NewInitializerChain(next)).Process()
}

func (b *ManagerBuilder) Load() *Index {
	return OpenIndex(b.directory)
}
