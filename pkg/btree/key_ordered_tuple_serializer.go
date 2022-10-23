package btree

import "github.com/oliversun9/chug/pkg/tuple"

type keyOrderedTupleSerializer struct{}

func (s keyOrderedTupleSerializer) Serialize(t tuple.Tuple) []byte {
	return nil
}

func (s keyOrderedTupleSerializer) Deserialize(b []byte, t tuple.Tuple) error {
	return nil
}
