package ralgebra

import "github.com/oliversun9/chug/pkg/tuple"

// Expression is a relational algebra relation.
type Expression interface {
	// Fetch loads value in the the map provided, and
	// returns true if the last one is read.
	// After the last value is fetched, Fetch will restart from the first one
	Fetch(map[string]tuple.Value) (bool, error)

	IsEmpty() bool
}
