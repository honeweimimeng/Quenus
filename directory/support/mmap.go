package support

import (
	"github.com/alexeymaximov/mmap"
	"os"
)

type MMapCore struct {
	File *os.File
	fMap *mmap.Mapping
	Len  int
}

func (r *MMapCore) Issued(file *os.File) (*MMapCore, error) {
	info, _ := file.Stat()
	fSize := uintptr(info.Size())
	size := uintptr(r.Len)
	if fSize > size {
		size = fSize
	}
	m, err := mmap.New(file.Fd(), 0, size, mmap.ModeReadWrite, 0)
	if err != nil {
		return nil, err
	}
	core := &MMapCore{File: file, fMap: m, Len: int(m.Length())}
	return core, nil
}

func (r *MMapCore) Read(p []byte) (n int, err error) {
	return r.ReadAt(p, 0)
}

func (r *MMapCore) ReadAt(p []byte, off int64) (n int, err error) {
	return r.fMap.ReadAt(p, off)
}

func (r *MMapCore) Write(p []byte) (n int, err error) {
	return r.WriteAt(p, 0)
}

func (r *MMapCore) WriteAt(p []byte, off int64) (n int, err error) {
	return r.fMap.WriteAt(p, off)
}

func (r *MMapCore) Close() error {
	return r.fMap.Close()
}
