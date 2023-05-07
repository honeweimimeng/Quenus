package directory

import (
	"errors"
	"github.com/honewemimeng/quenus/directory/support"
	"os"
)

const InitMapSize = 4 * 1024

type MMapDirectory struct {
	ctx  *Context
	core map[string]*support.MMapCore
}

func (d *MMapDirectory) Reader(file *FileSlice) DataIO {
	return d.core[file.Path]
}

func (d *MMapDirectory) Writer(file *FileSlice) DataIO {
	return d.core[file.Path]
}

func (d *MMapDirectory) Delete(file *FileSlice) {
	delete(d.core, file.Path)
}

func (d *MMapDirectory) FileOp(path string, cap int) *FileSlice {
	core := d.doMap(path, cap)
	slice := &FileSlice{File: core.File, Path: path}
	d.core[slice.Path] = core
	return slice
}

func (d *MMapDirectory) Exists(path string, create bool) (bool, error) {
	f, _ := d.openFile(path, create)
	if f == nil {
		return false, errors.New("file not found" + path)
	}
	return f != nil, nil
}

func (d *MMapDirectory) doMap(path string, cap int) *support.MMapCore {
	file, _ := d.openFile(path, false)
	if cap < 0 {
		info, _ := file.Stat()
		cap = int(info.Size())
	}
	m, err := support.IssuedMMap(file, cap)
	if err != nil {
		d.ctx.Log.Fatalln("cannot do mmap", err.Error())
	}
	return m
}

func (d *MMapDirectory) openFile(path string, create bool) (*os.File, os.FileInfo) {
	flag := os.O_RDWR
	if create {
		flag = os.O_RDWR | os.O_CREATE
	}
	f, err := os.OpenFile(path, flag, 0777)
	if err != nil {
		file := d.maybeCreateFile(err, path, create)
		if file != nil {
			info, _ := file.Stat()
			return file, info
		}
		return nil, nil
	}
	info, _ := f.Stat()
	return f, info
}

func (d *MMapDirectory) maybeCreateFile(err error, path string, create bool) *os.File {
	if os.IsNotExist(err) && create {
		file, err := os.Create(path)
		if err != nil {
			panic(err.Error())
		}
		return file
	}
	return nil
}
