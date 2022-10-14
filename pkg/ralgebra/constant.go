package ralgebra

import "github.com/oliversun9/chug/pkg/tuple"

// This is probably not needed:
// Cross product with a constant can be avoided by letting
// a predicate capture the constant's value
type Constant struct {
	Value tuple.Value
}
