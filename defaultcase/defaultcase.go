package defaultcase

var toDefaultCase func(name string) string = none

func ToDefaultCase(name string) string {
	return toDefaultCase(name)
}

func SetDefaultCase(c Case) {
	toDefaultCase = c.method()
}
