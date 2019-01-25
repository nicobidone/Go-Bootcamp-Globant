package selecteds

// PostgreSQL empty type
type PostgreSQL struct {
}

// Pused Option selected.
func Pused() string {
	return "You used PostgreSQL DB"
}

// Decir imprime rta
func (p PostgreSQL) Decir() string {
	return "do for PostgreSQL"
}
