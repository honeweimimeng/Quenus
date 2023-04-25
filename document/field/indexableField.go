package field

type IndexField interface {
	Id() int
	Name() string
	Val() any
	ValBin() []byte
}

type Proto struct {
	id   int
	name string
	val  any
}

func (p *Proto) Id() int {
	return p.id
}

func (p *Proto) Name() string {
	return p.name
}

func (p *Proto) Val() any {
	return p.val
}
func (p *Proto) ValBin() []byte {
	return nil
}
