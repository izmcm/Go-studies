package impl

import ()

type Calculadora struct{}

func (Calculadora) Add(p1 int, p2 int) int {
	return p1 + p2
}

func (Calculadora) Sub(p1 int, p2 int) int {
	return p1 - p2
}

func (Calculadora) Mul(p1 int, p2 int) int {
	return p1 * p2
}

func (Calculadora) Div(p1 int, p2 int) int {
	return p1 / p2
}
