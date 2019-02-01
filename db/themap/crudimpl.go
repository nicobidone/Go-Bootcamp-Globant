package themap

import "github.com/nicob/db/thedata"

// CrudImpl is a basic map structure
type CrudImpl struct {
	m map[string]thedata.Container
}

// NewMap map constructor
func NewMap() *CrudImpl {
	x := new(CrudImpl)
	x.m = map[string]thedata.Container{}
	return x
}

//Create bla bla
func (t CrudImpl) Create(key string, val thedata.Container) {
	_, ok := t.m[key]
	if !ok {
		t.m[key] = val
	}
}

// Read blabla
func (t CrudImpl) Read(key string) thedata.Container {
	return t.m[key]
}

//Update blabla
func (t CrudImpl) Update(key string, val thedata.Container) {
	t.m[key] = val
}

//Delete bla bla
func (t CrudImpl) Delete(key string) {
	delete(t.m, key)
}
