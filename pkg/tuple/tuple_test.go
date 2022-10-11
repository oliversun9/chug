package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTuple(t *testing.T) {
	t.Run("NewSchema test", testNewSchema)
	t.Run("NewTuple test", testNewTuple)
	t.Run("Schema.ValidateTuple test", testSchemaValidateTuple)
}

func testNewSchema(t *testing.T) {
	// Cannot be empty
	types := map[string]ValueType{}
	s, err := NewSchema(types)
	assert.Error(t, err)
	assert.Nil(t, s)

	// Int should work
	types = map[string]ValueType{
		"id": IntegerType,
	}
	s, err = NewSchema(types)
	assert.NoError(t, err)
	assert.Equal(t, IntegerType, s["id"])

	// String should work
	types = map[string]ValueType{
		"id": StringType,
	}
	s, err = NewSchema(types)
	assert.NoError(t, err)
	assert.Equal(t, StringType, s["id"])

	// Field name should allow letters in both cases and numbers
	types = map[string]ValueType{
		"id":  StringType,
		"Id1": IntegerType,
	}
	s, err = NewSchema(types)
	assert.NoError(t, err)
	assert.Equal(t, StringType, s["id"])
	assert.Equal(t, IntegerType, s["Id1"])

	// Empty name not allowed
	types = map[string]ValueType{
		"": StringType,
	}
	s, err = NewSchema(types)
	assert.Error(t, err)
	assert.Nil(t, s)

	// Other characters not allowed in column name
	types = map[string]ValueType{
		"-": StringType,
	}
	s, err = NewSchema(types)
	assert.Error(t, err)
	assert.Nil(t, s)
}

func testNewTuple(t *testing.T) {
	types := map[string]ValueType{
		"id":      IntegerType,
		"name":    StringType,
		"goals":   IntegerType,
		"assists": IntegerType,
		"age":     IntegerType,
	}
	s, err := NewSchema(types)
	assert.NoError(t, err)

	values1 := map[string]Value{
		"id":     IntegerValue(5),
		"name":   StringValue("Dennis Begkamp"),
		"goals":  IntegerValue(999),
		"assits": IntegerValue(999), // WRONG COLUMN NAME
		"age":    IntegerValue(50),
	}
	tuple, err := NewTuple(s, values1)
	assert.Error(t, err)
	assert.Nil(t, tuple)

	// Missing  a field
	values2 := map[string]Value{
		"id":    IntegerValue(5),
		"name":  StringValue("Dennis Begkamp"),
		"goals": IntegerValue(999),
		"age":   IntegerValue(50),
	}
	tuple, err = NewTuple(s, values2)
	assert.Error(t, err)
	assert.Nil(t, tuple)

	// Good
	values3 := map[string]Value{
		"assists": IntegerValue(1000),
		"id":      IntegerValue(5),
		"name":    StringValue("Dennis Begkamp"),
		"goals":   IntegerValue(999),
		"age":     IntegerValue(50),
	}
	tuple, err = NewTuple(s, values3)
	assert.NoError(t, err)
	assert.Equal(t, len(values3), len(tuple))
	for k, v := range values3 {
		assert.Equal(t, v, tuple[k])
	}
}

func testSchemaValidateTuple(t *testing.T) {
	s, err := NewSchema(
		map[string]ValueType{
			"id":      IntegerType,
			"name":    StringType,
			"goals":   IntegerType,
			"assists": IntegerType,
			"age":     IntegerType,
		},
	)
	assert.NoError(t, err)
	tuple, err := NewTuple(
		s,
		map[string]Value{
			"assists": IntegerValue(1000),
			"id":      IntegerValue(5),
			"name":    StringValue("Dennis Begkamp"),
			"goals":   IntegerValue(999),
			"age":     IntegerValue(50),
		},
	)
	assert.NoError(t, err)

	t.Run("The schema that initialized the tuple should be OK", func(t *testing.T) {
		err = s.ValidateTuple(tuple)
		assert.NoError(t, err)
	})

	t.Run("Different field name should error", func(t *testing.T) {
		s, err := NewSchema(
			map[string]ValueType{
				"ID":      IntegerType,
				"name":    StringType,
				"goals":   IntegerType,
				"assists": IntegerType,
				"age":     IntegerType,
			},
		)
		assert.NoError(t, err)
		err = s.ValidateTuple(tuple)
		assert.Error(t, err)
	})

	t.Run("Different length should error", func(t *testing.T) {
		s, err := NewSchema(
			map[string]ValueType{
				"id":    IntegerType,
				"name":  StringType,
				"goals": IntegerType,
				"age":   IntegerType,
			},
		)
		assert.NoError(t, err)
		err = s.ValidateTuple(tuple)
		assert.Error(t, err)
	})

	t.Run("Different type should error", func(t *testing.T) {
		s, err := NewSchema(
			map[string]ValueType{
				"id":      IntegerType,
				"name":    StringType,
				"goals":   StringType,
				"assists": IntegerType,
				"age":     IntegerType,
			},
		)
		assert.NoError(t, err)
		err = s.ValidateTuple(tuple)
		assert.Error(t, err)
	})

	t.Run("Different order should be OK", func(t *testing.T) {
		s, err := NewSchema(
			map[string]ValueType{
				"name":    StringType,
				"goals":   IntegerType,
				"assists": IntegerType,
				"age":     IntegerType,
				"id":      IntegerType,
			},
		)
		assert.NoError(t, err)
		err = s.ValidateTuple(tuple)
		assert.NoError(t, err)
	})
}
