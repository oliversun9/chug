package table

import (
	"fmt"

	"github.com/oliversun9/chug/pkg/tuple"
)

type Table interface {
	// Table's name
	Name() string

	// Shows the schema
	Schema() tuple.Schema

	// True if table has no records
	IsEmpty() bool

	// Create a record.
	Create(tuple.Tuple) error

	// returns the tuple read and whether it's the last one
	Read() (tuple.Tuple, bool, error)

	// TODO: add something like seek?

	// Update the record with the same primary key.
	Update(tuple.Tuple) error

	// Delete the record with primary key
	Delete(tuple.Tuple) error
}

func CreateTable(string, tuple.Schema) (Table, error) {
	return nil, fmt.Errorf("unimplemented")
}

func LoadTable(string) (Table, error) {
	return nil, fmt.Errorf("unimplemented")
}
