package tipos

// InternalError contains
type InternalError struct {
	Res string
}

// NewThing inic
func NewThing() *InternalError {
	p := new(InternalError)
	p.Res = "Error Interno" // <- a very sensible default value
	return p
}

func (i *InternalError) Error() string {

	return i.Res
}
