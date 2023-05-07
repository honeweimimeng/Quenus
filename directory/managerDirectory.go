package directory

import (
	"fmt"
	"github.com/honewemimeng/quenus/directory/support"
)

type ManagerDirectory struct {
	ctx      *Context
	coreDire Directory
}

func NewManagerDirectory(root string) Directory {
	ctx := NewDirectoryContext(root)
	mMap := &MMapDirectory{ctx: ctx, core: make(map[string]*support.MMapCore)}
	return &ManagerDirectory{ctx: ctx, coreDire: mMap}
}

func (d *ManagerDirectory) Reader(file *FileSlice) DataIO {
	return d.coreDire.Reader(file)
}

func (d *ManagerDirectory) Writer(file *FileSlice) DataIO {
	return d.coreDire.Writer(file)
}

func (d *ManagerDirectory) Delete(file *FileSlice) {
	d.coreDire.Delete(file)
}

func (d *ManagerDirectory) FileOp(path string, cap int) *FileSlice {
	s := d.NormalizePath(FileSlice{Path: path})
	return d.coreDire.FileOp(s.Path, cap)
}

func (d *ManagerDirectory) Exists(path string, create bool) (bool, error) {
	s := d.NormalizePath(FileSlice{Path: path})
	return d.coreDire.Exists(s.Path, create)
}

// NormalizePath Why Not Return *FileSlice ?
// copy In stack can avoid move FileSlice to heap
func (d *ManagerDirectory) NormalizePath(file FileSlice) FileSlice {
	file.Path = fmt.Sprintf("%s/%s", d.ctx.RootPath, file.Path)
	return file
}
