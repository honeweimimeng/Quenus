package directory

import "io"

type Directory interface {
	ReadString() string
	ReadBytes(bytes *[]byte)
	Reader() io.Reader
	Writer() io.Writer
	Close() io.Closer
}
