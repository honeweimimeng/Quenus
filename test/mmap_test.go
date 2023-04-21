package test

import (
	"github.com/alexeymaximov/mmap"
	"os"
	"path/filepath"
	"testing"
)

var testData = []byte("0123456789ABCDEF")
var testPath = filepath.Join("C:\\workspace\\quenus\\", "data\\temp.txt")

func init() {
	f := openFile(os.O_RDWR | os.O_CREATE | os.O_TRUNC)
	f.Write(testData)
	f.Close()
}

func openFile(flags int) *os.File {
	f, err := os.OpenFile(testPath, flags, 0644)
	if err != nil {
		panic(err.Error())
	}
	return f
}

var testBuffer = []byte{'b', 'E', 'L', 'L', 'O'}

func TestUnmap(t *testing.T) {
	f := openFile(os.O_RDWR)
	defer f.Close()
	mapping, err := mmap.New(f.Fd(), 0, 1024, mmap.ModeReadWrite, 0)
	if _, err := mapping.WriteAt(testBuffer, 0); err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Errorf("error mapping: %s", err)
	}
	if err := mapping.Close(); err != nil {
		t.Errorf("error unmapping: %s", err)
	}
}
