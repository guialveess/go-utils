package math

var Hello string = "Hello, Math!"

type Math struct {
	A int
	B int
}

var result int

func (m *Math) Add() int {
	result = m.A + m.B
	return result
}

func (m *Math) Subtract() int {
	return m.A - m.B
}

func (m *Math) Multiply() int {
	return m.A * m.B
}

func (m *Math) Divide() int {
	return m.A / m.B
}
