package directory

import (
	"io"
	"os"
)

type DataIO interface {
	ReadAll() ([]byte, error)
	io.Writer
	io.WriterAt
	io.Reader
	io.ReaderAt
	io.Closer
}

type FileSlice struct {
	File *os.File
	Path string `json:"path"`
	Cur  uint64 `json:"cur"`
}

func (f *FileSlice) Close() error {
	return f.File.Close()
}

type Directory interface {
	FileOp(path string, cap int) *FileSlice
	Reader(file *FileSlice) DataIO
	Writer(file *FileSlice) DataIO
	Exists(path string, create bool) (bool, error)
	Delete(file *FileSlice)
}

type IOErr struct {
	Err string
}

func (e *IOErr) Error() string {
	return e.Err
}

/*
* async
 */

type DataCallable func(io DataIO)

type AsyncDirectory interface {
	Directory
	AsyncCall(call DataCallable)
}
