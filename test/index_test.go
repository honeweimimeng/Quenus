package test

import (
	"github.com/honewemimeng/quenus/directory"
	"github.com/honewemimeng/quenus/document"
	"github.com/honewemimeng/quenus/index"
	"testing"
	"time"
)

func TestAddDoc(t *testing.T) {
}

func TestIndex(t *testing.T) {
	d := directory.NewManagerDirectory("C:\\workspace\\quenus\\data")
	manager := index.NewManagerBuilder().Build(d).StartListener()
	manager.Write(&document.Document{}, &document.Document{})
	time.Sleep(200 * time.Second)
	println(manager)
}
