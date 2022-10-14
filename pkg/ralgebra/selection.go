package ralgebra

type Selection struct {
	E Expression
	P Predicate
}

// Predicate is used by Select
type Predicate func(Expression) bool
