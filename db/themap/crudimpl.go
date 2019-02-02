package themap

import (
	"errors"
	"fmt"

	"github.com/nicob/db/thedata"
)

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
func (t CrudImpl) Read(key string) (thedata.Container, error) {
	value, ok := t.m[key]
	if !ok {
		x := *thedata.NewContainer("", "")
		return x, errors.New("The value is not avaible")
	}
	return value, nil
}

//Update blabla
func (t CrudImpl) Update(key string, val thedata.Container) {
	t.m[key] = val
}

//Delete bla bla
func (t CrudImpl) Delete(key string) {
	delete(t.m, key)
}

func (t CrudImpl) Error() string {
	return fmt.Sprint("El valor no se encuentra disponible")
}
