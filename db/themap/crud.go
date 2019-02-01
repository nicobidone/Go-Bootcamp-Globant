package themap

import "github.com/nicob/db/thedata"

// Crud CRUD
type Crud interface {
	Create(key string, val thedata.Container)
	Read(key string) string
	Update(key string, val thedata.Container)
	Delete(key string)
}
