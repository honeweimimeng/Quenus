package index

import "github.com/honeweimimeng/eventgo/driver/event"

type ManagerInitializer interface {
	Process(m *Manager) bool
	Next() ManagerInitializer
}

type InitializerChain struct {
	manager *Manager
	next    ManagerInitializer
}

func NewInitializerChain(next ManagerInitializer) *InitializerChain {
	return &InitializerChain{next: next}
}

func (c *InitializerChain) ProcessChain(m *Manager, init ManagerInitializer) {
	if init != nil && init.Process(m) {
		c.ProcessChain(m, c.Next())
	}
}

func (c *InitializerChain) Next() ManagerInitializer {
	return c.next
}

func (c *InitializerChain) Process(m *Manager) bool {
	c.manager = m
	c.AddEventHandle(NewReader(), func(ch chan []event.Proto) {
		c.manager.readAdvice = ch
	}).AddEventHandle(NewWriter(), func(ch chan []event.Proto) {
		c.manager.writeAdvice = ch
	})
	return c.Next() != nil
}

func (c *InitializerChain) AddEventHandle(eventR Registry, setter event.SimpleTrigger) *InitializerChain {
	c.manager.boot.EventLoop().Handle(eventR.EventHandle()).ExTrigger(setter)
	return c
}
