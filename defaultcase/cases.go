package defaultcase

import "github.com/iancoleman/strcase"

const (
	None Case = iota
	Snak_case
	CableCase
)

type Case int

func (c Case) method() func(string) string {
	switch c {
	default:
		return none
	case CableCase:
		return cableCase
	case Snak_case:
		return snak_case
	}
}

func none(n string) string {
	return n
}

func cableCase(name string) string {
	defaultName := strcase.ToLowerCamel(name)
	return defaultName
}

func snak_case(name string) string {
	defaultName := strcase.ToSnake(name)
	return defaultName
}
