package directory

import "github.com/sirupsen/logrus"

type Context struct {
	Log      *logrus.Logger
	RootPath string
}

func NewDirectoryContext(root string) *Context {
	return &Context{RootPath: root, Log: logrus.New()}
}
