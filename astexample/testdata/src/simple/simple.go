package packagesimple

import "errors"

const (
	myconstant = "myconstant"
	MyConstant = 1
)

type MyStruct struct {
	a, b int
}

func MustNotZero(a, b int) (MyStruct, error) {
	m := MyStruct{a: a, b: b}
	if m.IsZero() {
		return m, errors.New("mystruct can't be zero")
	}
	return m, nil
}

func (m MyStruct) IsZero() bool {
	return m.a == 0 && m.b == 0
}
