package tipos

// ThirdPartyError struct
type ThirdPartyError struct {
	Res string
}

// NewThirdPartyError inic
func NewThirdPartyError() *ThirdPartyError {
	p := new(ThirdPartyError)
	p.Res = "Error de terceros" // <- a very sensible default value
	return p
}

func (i ThirdPartyError) Error() string {
	return i.Res
}
