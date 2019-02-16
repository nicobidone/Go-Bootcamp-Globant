package entity

type Cart struct {
	id          int32
	description string
	content     []CartContent
}

func NewCart(id int32, customer string) *Cart {
	x := new(Cart)
	x.id = id
	x.description = ""
	x.content = []CartContent{}
	return x
}

func (c *Cart) AddCartContent(cartcontent Product, quantity int) {
	p := NewCartContent(cartcontent, quantity)
	c.content = append(c.content, *p)
}

func (c *Cart) GetCartContent() []CartContent {
	return c.content
}

func (c *Cart) positionCartContent(id int) int {
	for i := 0; i < len(c.content); i++ {
		if c.content[i].IsContent(id) {
			return i
		}
	}
	return 0
}

func (c *Cart) EditCartContent(id int, quantity int) {
	i := c.positionCartContent(id)

	c.content[i].SetQuantityCartContent(quantity)
}

func (c *Cart) RemoveCartContent(id int) {
	i := c.positionCartContent(id)

	copy(c.content[i:], c.content[i+1:])
	c.content[len(c.content)-1] = BlankCartContent()
	c.content = c.content[:len(c.content)-1]
}

func (c *Cart) ClearCartContent() {
	c.content = []CartContent{}
}
