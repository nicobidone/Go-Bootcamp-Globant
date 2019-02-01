package thedata

type Container struct {
	a string
	b string
}

func NewContainer(a string, b string) *Container {
	x := new(Container)
	x.a = a
	x.b = b
	return x
}

func (c Container) AValue() string {
	return c.a
}

func (c Container) BValue() string {
	return c.b
}
