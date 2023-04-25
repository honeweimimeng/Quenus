package test

import (
	"github.com/honewemimeng/quenus/directory"
	"github.com/honewemimeng/quenus/document"
	"github.com/honewemimeng/quenus/index"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestAddDoc(t *testing.T) {
}

func TestIndex(t *testing.T) {
	ctx := &directory.Context{Log: logrus.New()}
	d := directory.NewMMapDirectory(ctx)
	manager := index.NewManagerBuilder(d).Build().StartListener()
	manager.Write(&document.Document{}, &document.Document{})
	time.Sleep(200 * time.Second)
	println(manager)
}

func TestPageCache(t *testing.T) {
}
