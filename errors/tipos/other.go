package tipos

// OtherError struct
type OtherError struct {
	Res string
}

// NewOtherError inic
func NewOtherError() *OtherError {
	p := new(OtherError)
	p.Res = "Error de otro tipo" // <- a very sensible default value
	return p
}

func (i OtherError) Error() string {
	return i.Res
}
