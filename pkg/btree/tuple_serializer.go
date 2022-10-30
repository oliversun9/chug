package btree

import "github.com/oliversun9/chug/pkg/tuple"

// Strategy for serialize and deserialize tuples,
// without passing in a schema each time.
type TupleSerializer interface {
	Serialize(tuple.Tuple) []byte
	Deserialize([]byte, *tuple.Tuple) error
}
