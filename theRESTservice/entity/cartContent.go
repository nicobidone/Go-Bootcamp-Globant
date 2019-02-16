package entity

import "strconv"

type CartContent struct {
	p Product
	q int
}

func NewCartContent(prod Product, quantity int) *CartContent {
	x := new(CartContent)
	x.p = prod
	x.q = quantity
	return x
}

func (c *CartContent) IsContent(id int) bool {
	if c.p.Id == strconv.Itoa(id) {
		return true
	}
	return false
}

func BlankCartContent() CartContent {
	return *new(CartContent)
}

func (c *CartContent) SetQuantityCartContent(quantity int) {
	c.q = quantity
}
