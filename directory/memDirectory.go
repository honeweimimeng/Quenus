package directory

import (
	"github.com/honewemimeng/quenus/directory/support"
	"io"
)

type MemDirectory struct {
	core support.MMapCore
}

func (d *MemDirectory) Reader() io.Reader {
	return &d.core
}

func (d *MemDirectory) Writer() io.Writer {
	return &d.core
}

func (d *MemDirectory) Close() io.Closer {
	return nil
}

func (d *MemDirectory) ReadString() string {
	return ""
}

func (d *MemDirectory) ReadBytes(bytes *[]byte) {

}
