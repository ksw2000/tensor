package internal

import "golang.org/x/exp/constraints"

type Addable interface {
	constraints.Float | constraints.Integer | constraints.Complex | ~string
}

type Num interface {
	constraints.Float | constraints.Integer | constraints.Complex
}

type OrderedNum interface {
	constraints.Float | constraints.Integer
}

// Applicable is any function that may be applied
type Applicable[DT any] interface {
	func(DT) DT | func(DT) (DT, error)
}
