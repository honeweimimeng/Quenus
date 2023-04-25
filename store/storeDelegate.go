package store

type Delegate struct {
}

func (d *Delegate) NewDelegate() *Delegate {
	return &Delegate{}
}
