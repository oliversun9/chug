package tuple

import (
	"fmt"
)

type Tuple map[string]Value

func NewTuple(s Schema, values map[string]Value) (Tuple, error) {
	if len(s) != len(values) {
		return nil, fmt.Errorf("schema has arity %d, while %d values are provided", len(s), len(values))
	}

	t := make(map[string]Value, len(s))
	for col, val := range values {
		expectedType, ok := s[col]
		if !ok {
			return nil, fmt.Errorf("unexpectd Column name %s is not in the schema", col)
		}
		if expectedType != val.ValueType() {
			return nil, fmt.Errorf("types don't match, expected %v, received %v", expectedType, val.ValueType())
		}
		t[col] = val
	}
	return t, nil
}
