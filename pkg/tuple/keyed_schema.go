package tuple

import "fmt"

type KeyedSchema interface {
	Schema() Schema
	KeyName() string
	KeyType() ValueType
}

func NewKeyedSchema(keyName string, keyType ValueType, m map[string]ValueType) (KeyedSchema, error) {
	typeFromMap, ok := m[keyName]
	if !ok {
		return nil, fmt.Errorf("key name '%s' is not among the column names provided", keyName)
	}
	if typeFromMap != keyType {
		return nil, fmt.Errorf("expected type %v from map, got %v instead", keyType, typeFromMap)
	}
	schema, err := newSchema(m)
	if err != nil {
		return nil, err
	}
	return keyedSchema{
		schema:  schema,
		keyName: keyName,
		keyType: keyType,
	}, nil
}

type keyedSchema struct {
	schema  Schema
	keyName string
	keyType ValueType
}

func (ks keyedSchema) Schema() Schema {
	return ks.schema
}

func (ks keyedSchema) KeyName() string {
	return ks.keyName
}
func (ks keyedSchema) KeyType() ValueType {
	return ks.keyType
}
