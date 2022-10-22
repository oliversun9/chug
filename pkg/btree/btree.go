package btree

import "github.com/oliversun9/chug/pkg/tuple"

type BTree struct {
}

// Table's name
func (b *BTree) Name() string {
	return ""
}

// Shows the schema
func (b *BTree) Schema() tuple.Schema {
	return nil
}

// True if table has no records
func (b *BTree) IsEmpty() bool {
	return false
}

// Create a record.
func (b *BTree) Create(tuple.Tuple) error {
	return nil
}

// returns the tuple read and whether it's the last one
func (b *BTree) Read() (tuple.Tuple, bool, error) {
	return nil, false, nil
}

// TODO: add something like seek?

// Update the record with the same primary key.
func (b *BTree) Update(tuple.Tuple) error {
	return nil
}

// Delete the record with primary key
func (b *BTree) Delete(tuple.Tuple) error {
	return nil
}
