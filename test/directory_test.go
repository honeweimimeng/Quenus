package test

import (
	"github.com/honewemimeng/quenus/directory"
	"testing"
)

func TestMMap(t *testing.T) {
	d := &directory.MMapDirectory{}
	path := "C:\\workspace\\quenus\\data\\temp.txt"
	res := make([]byte, 10)
	fSlice := d.FileOp(path)
	defer func() {
		err := fSlice.Close()
		if err != nil {
			panic(err.Error())
		}
	}()
	i := fSlice.Io
	_, err := i.Write([]byte{'c', 'e', 'l', 'l', 'd'})
	if err != nil {
		panic(err.Error())
	}
	_, err = i.Read(res)
	println(string(res))
}

func TestMem(t *testing.T) {

}

func TestFeedCache(t *testing.T) {

}
