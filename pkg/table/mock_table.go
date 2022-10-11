package table

import (
	"fmt"

	"github.com/oliversun9/chug/pkg/tuple"
)

type MockTable struct {
	name    string
	schema  tuple.Schema
	records []tuple.Tuple
	cursor  int
}

func NewMockTable(name string, schema tuple.Schema, tuples []tuple.Tuple) (Table, error) {
	for _, tuple := range tuples {
		if err := schema.ValidateTuple(tuple); err != nil {
			return nil, err
		}
	}
	return &MockTable{
		name:    name,
		schema:  schema,
		records: tuples,
		cursor:  0,
	}, nil
}

func (t *MockTable) Name() string {
	return t.name
}

func (t *MockTable) Schema() tuple.Schema {
	return t.schema
}

func (t *MockTable) IsEmpty() bool {
	return len(t.records) == 0
}

func (t *MockTable) Create(tp tuple.Tuple) error {
	t.records = append(t.records, tp)
	return nil
}

func (t *MockTable) Read() (tuple.Tuple, bool, error) {
	if !(0 <= t.cursor && t.cursor < len(t.records)) {
		return nil, false, fmt.Errorf("bug in MockTable, misplaced cursor of value %d", t.cursor)
	}
	tp := t.records[t.cursor]
	isLast := false
	if t.cursor++; t.cursor >= len(t.records) {
		t.cursor = 0
		isLast = true
	}
	return tp, isLast, nil
}

func (t *MockTable) Update(tuple.Tuple) error { return nil }

func (t *MockTable) Delete(tuple.Tuple) error { return nil }
