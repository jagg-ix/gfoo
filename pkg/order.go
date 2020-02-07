package gfoo

type Order = int

const (
	Lt = Order(-1)
	Eq = Order(0)
	Gt = Order(1)
)

func CompareInt(x, y int) Order {
	if x < y {
		return Lt
	}

	if x > y {
		return Gt
	}

	return Eq
}

