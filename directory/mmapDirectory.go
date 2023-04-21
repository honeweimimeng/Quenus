package directory

import (
	"errors"
	"github.com/honewemimeng/quenus/directory/support"
	"os"
)

type MMapDirectory struct {
	ctx  *Context
	core support.MMapCore
}

func NewMMapDirectory(ctx *Context) Directory {
	return &MMapDirectory{ctx: ctx}
}

func (d *MMapDirectory) Read(path string) DataIO {
	return d.doMap(path)
}

func (d *MMapDirectory) Write(path string) DataIO {
	return d.doMap(path)
}

func (d *MMapDirectory) Delete(path string) {

}

func (d *MMapDirectory) FileOp(path string) *FileSlice {
	core := d.doMap(path)
	slice := &FileSlice{File: core.File, Path: path, Io: core}
	return slice
}

func (d *MMapDirectory) Exists(path string) (bool, error) {
	f, _ := d.openFile(path)
	if f == nil {
		return false, errors.New("file not found" + path)
	}
	return true, nil
}

func (d *MMapDirectory) doMap(path string) *support.MMapCore {
	file, _ := d.openFile(path)
	m, err := d.core.Issued(file)
	if err != nil {
		d.ctx.Log.Fatalln("cannot do mmap", err.Error())
	}
	return m
}

func (d *MMapDirectory) openFile(path string) (*os.File, os.FileInfo) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		file := d.maybeCreateFile(err, path)
		if file != nil {
			info, _ := file.Stat()
			return file, info
		}
		d.ctx.Log.Fatalln("open file fail,", err.Error())
		return nil, nil
	}
	info, _ := f.Stat()
	return f, info
}

func (d *MMapDirectory) maybeCreateFile(err error, path string) *os.File {
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			panic(err.Error())
		}
		return file
	}
	return nil
}
