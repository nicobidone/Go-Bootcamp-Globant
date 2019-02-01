package themap

import "github.com/nicob/db/thedata"

// Themap is a basic map structure
type crud_impl struct {
	m map[string]thedata.Container
}

// NewMap map constructor
func NewMap() *crud_impl {
	x := new(crud_impl)
	x.m = map[string]thedata.Container{}
	return x
}

//Create bla bla
func (t crud_impl) Create(key string, val thedata.Container) {
	_, ok := t.m[key]
	if !ok {
		t.m[key] = val
	}
}

// Read blabla
func (t crud_impl) Read(key string) thedata.Container {
	return t.m[key]
}

//Update blabla
func (t crud_impl) Update(key string, val thedata.Container) {
	t.m[key] = val
}

//Delete bla bla
func (t crud_impl) Delete(key string) {
	delete(t.m, key)
}
