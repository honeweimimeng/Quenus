package directory

import (
	"io"
	"os"
)

type DataIO interface {
	io.Writer
	io.WriterAt
	io.Reader
	io.ReaderAt
	io.Closer
}

type FileSlice struct {
	File *os.File
	Path string
	Io   DataIO
}

func (f *FileSlice) Close() error {
	err := f.Io.Close()
	if err != nil {
		return err
	}
	return f.File.Close()
}

type Directory interface {
	FileOp(path string) *FileSlice
	Read(path string) DataIO
	Write(path string) DataIO
	Exists(path string) (bool, error)
	Delete(path string)
}

/*
* async
 */

type DataCallable func(io DataIO)

type AsyncDirectory interface {
	Directory
	AsyncCall(call DataCallable)
}
