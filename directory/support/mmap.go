package support

import (
	"github.com/alexeymaximov/mmap"
	"os"
)

type MMapCore struct {
	File *os.File
	fMap *mmap.Mapping
	Cap  int
}

func IssuedMMap(file *os.File, cap int) (*MMapCore, error) {
	core := &MMapCore{File: file}
	return core, core.getMMap(file, cap)
}

func (r *MMapCore) getMMap(file *os.File, cap int) error {
	fSize := uintptr(cap)
	if fSize == 0 {
		r.fMap = &mmap.Mapping{}
		return nil
	}
	m, err := mmap.New(file.Fd(), 0, fSize, mmap.ModeReadWrite, 0)
	if err != nil {
		return err
	}
	r.fMap = m
	r.Cap = cap
	return nil
}

func (r *MMapCore) Read(p []byte) (n int, err error) {
	return r.ReadAt(p, 0)
}

func (r *MMapCore) ReadAt(p []byte, off int64) (n int, err error) {
	return r.fMap.ReadAt(p, off)
}

func (r *MMapCore) Write(p []byte) (n int, err error) {
	info, _ := r.File.Stat()
	return r.WriteAt(p, info.Size())
}

func (r *MMapCore) WriteAt(p []byte, off int64) (int, error) {
	_ = r.fMap.Lock()
	defer func() { _ = r.fMap.Unlock() }()
	newSize := int(off) + len(p)
	if newSize > int(r.fMap.Length()) {
		err := r.refreshMMap(newSize, int(r.fMap.Length()))
		if err != nil {
			return 0, err
		}
	}
	return r.fMap.WriteAt(p, off)
}

func (r *MMapCore) ReadAll() ([]byte, error) {
	info, _ := r.File.Stat()
	data := make([]byte, info.Size())
	_, err := r.Read(data)
	return data, err
}

func (r *MMapCore) refreshMMap(newSize int, oldSize int) error {
	added := make([]byte, newSize-oldSize)
	_, err := r.File.Write(added)
	if err != nil {
		return err
	}
	err = r.getMMap(r.File, newSize)
	if err != nil {
		return err
	}
	return nil
}

func (r *MMapCore) CurMem() []byte {
	return nil
}

func (r *MMapCore) Close() error {
	return r.fMap.Close()
}
