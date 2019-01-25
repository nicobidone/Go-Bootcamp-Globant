package selecteds

// MongoDB empty struct
type MongoDB struct {
}

// Mused Operation Selected
func Mused() string {
	return "You used MongoDB DB"
}

// Decir imprime rta
func (m MongoDB) Decir() string {
	return "do for MongoDB"
}
