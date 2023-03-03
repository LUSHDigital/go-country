package country

//go:generate go run gen.go

type Country struct {
	Name   string
	Alpha2 string
	Alpha3 string
}
