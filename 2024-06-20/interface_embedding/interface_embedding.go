package interface_embedding

type Ie interface {
	embed
	SetId(id int)
}

type embed interface {
	SetName(name string)
}

type Putter struct {
	id   int
	name string
}

func (p *Putter) SetId(id int) {
	p.id = id
}

func (p *Putter) SetName(name string) {
	p.name = name
}
