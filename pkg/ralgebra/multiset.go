package ralgebra

import (
	"fmt"

	"github.com/oliversun9/chug/pkg/table"
	"github.com/oliversun9/chug/pkg/tuple"
)

type Multiset struct {
	// Name is used to distinguish between different tuples sharing common column names.
	// e.g. In R1 x R2, need to refer their ids as "r1.id" and "r2.id", r1 and r2 are names.
	Name  string
	Table table.Table
}

func (m Multiset) Fetch(dest map[string]tuple.Value) (bool, error) {
	if m.Table.IsEmpty() {
		return false, fmt.Errorf("unable to fetch from an empty table: %s", m.Table.Name())
	}
	tuple, isLast, err := m.Table.Read()
	if err != nil {
		return false, fmt.Errorf("error reading from table %s: %w", m.Table.Name(), err)
	}
	for col, val := range tuple {
		universalColumnName := fmt.Sprintf("%s.%s", m.Name, col)
		dest[universalColumnName] = val
	}
	return isLast, nil
}

func (m Multiset) IsEmpty() bool {
	return m.Table.IsEmpty()
}
