package struct_embedding

type A struct {
	B
	I
}

type I interface {
	GetName() string
}

type B struct {
	Id int
}

func (b B) GetId() int {
	return b.Id
}

type C struct {
	Name string
}

func (c C) GetName() string {
	return c.Name
}
