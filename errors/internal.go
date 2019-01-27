package main

// InternalError contains
type InternalError struct {
	Res string
}

func (i *InternalError) Error() string {

	return i.Res
}
