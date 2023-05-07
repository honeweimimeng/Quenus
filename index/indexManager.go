package index

import (
	"github.com/honeweimimeng/eventgo/driver/event"
	"github.com/honewemimeng/quenus/document"
)

type Registry interface {
	EventHandle() ([]event.Proto, event.SimpleHandler)
}

type Manager struct {
	ctx         *IContext
	idx         *Index
	writeAdvice chan []event.Proto
	readAdvice  chan []event.Proto
	boot        event.BootStrap
	InitChain   *InitializerChain
}

func NewManager(idx *Index, init ManagerInitializer) *Manager {
	manager := &Manager{
		idx:       idx,
		InitChain: &InitializerChain{next: init},
		ctx:       NewIndexCtx(),
	}
	return manager
}

func (m *Manager) StartListener() *Manager {
	m.boot.StartUp()
	return m
}

func (m *Manager) Write(doc ...*document.Document) {
	m.publishDoc(m.writeAdvice, func(doc *document.Document) event.Proto {
		return NewWRITE(doc)
	}, doc...)
}

func (m *Manager) Read(doc ...*document.Document) {
	m.publishDoc(m.readAdvice, func(doc *document.Document) event.Proto {
		return NewREAD(doc)
	}, doc...)
}

func (m *Manager) publishDoc(ch chan []event.Proto,
	instance func(doc *document.Document) event.Proto,
	doc ...*document.Document) {
	protoArr := make([]event.Proto, len(doc))
	//batch,avoid doc copy in range
	for i := 0; i < len(doc); i++ {
		protoArr[i] = instance(doc[i])
	}
	ch <- protoArr
}

func (m *Manager) Process() *Manager {
	m.InitChain.ProcessChain(m, m.InitChain)
	return m
}
