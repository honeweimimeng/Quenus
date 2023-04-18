package field

type IndexField interface {
	Name() string
	Val() any
	ValBin() []byte
}

type Proto struct {
	name string
	val  any
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
