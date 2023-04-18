package directory

import (
	"github.com/honewemimeng/quenus/directory/support"
	"io"
)

type MMapDirectory struct {
	core *support.MMapCore
}

func (d *MMapDirectory) Reader() io.Reader {
	return d.core
}

func (d *MMapDirectory) Writer() io.Writer {
	return d.core
}

func (d *MMapDirectory) Close() io.Closer {
	return nil
}

func (d *MMapDirectory) ReadString() string {
	return ""
}

func (d *MMapDirectory) ReadBytes(bytes *[]byte) {

}
