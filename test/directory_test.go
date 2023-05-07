package test

import (
	"github.com/honewemimeng/quenus/directory"
	"testing"
)

func TestMMap(t *testing.T) {
	d := &directory.MMapDirectory{}
	path := "C:\\workspace\\quenus\\data\\temp.txt"
	res := make([]byte, 10)
	fSlice := d.FileOp(path, -1)
	defer func() {
		err := fSlice.Close()
		if err != nil {
			panic(err.Error())
		}
	}()
	println(string(res))
}

func TestMem(t *testing.T) {

}

func TestFeedCache(t *testing.T) {

}
