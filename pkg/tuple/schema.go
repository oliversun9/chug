package tuple

import "fmt"

type Schema map[string]ValueType

func NewSchema(m map[string]ValueType) (Schema, error) {
	if len(m) <= 0 {
		return nil, fmt.Errorf("schema must not be empty")
	}
	types := make(map[string]ValueType, len(m))
	for name, valueType := range m {
		if !isValidColumnName(name) {
			return nil, fmt.Errorf("%s is not a valid column name", name)
		}
		types[name] = valueType
	}
	return types, nil
}

func isValidColumnName(name string) bool {
	if name == "" {
		return false
	}
	for _, c := range name {
		if 'a' <= c && c <= 'z' {
			continue
		} else if 'A' <= c && c <= 'Z' {
			continue
		} else if '0' <= c && c <= '9' {
			continue
		}
		return false
	}
	return true
}
