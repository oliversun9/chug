package ralgebra

import (
	"testing"

	"github.com/oliversun9/chug/pkg/algorithm"
	"github.com/oliversun9/chug/pkg/table"
	"github.com/oliversun9/chug/pkg/tuple"
	"github.com/stretchr/testify/assert"
)

func TestMultiset(t *testing.T) {
	types := map[string]tuple.ValueType{
		"id":   tuple.IntegerType,
		"name": tuple.StringType,
	}
	s, err := tuple.NewSchema(types)
	assert.NoError(t, err)

	data := []map[string]tuple.Value{
		{
			"id":   tuple.IntegerValue(1),
			"name": tuple.StringValue("Name A"),
		},
		{
			"id":   tuple.IntegerValue(2),
			"name": tuple.StringValue("Name B"),
		},
		{
			"id":   tuple.IntegerValue(3),
			"name": tuple.StringValue("Name C"),
		},
		{
			"id":   tuple.IntegerValue(4),
			"name": tuple.StringValue("Name D"),
		},
		{
			"id":   tuple.IntegerValue(5),
			"name": tuple.StringValue("Name E"),
		},
	}
	tuples, err := algorithm.Transform(
		data,
		func(values map[string]tuple.Value) (tuple.Tuple, error) {
			return tuple.NewTuple(s, values)
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, tuples)
	mt, err := table.NewMockTable(
		"mt1",
		s,
		tuples,
	)
	assert.NoError(t, err)

	for i, expected := range tuples {
		tableTuple, isLast, err := mt.Read()
		assert.NoError(t, err)
		if i == len(tuples)-1 {
			assert.Equal(t, true, isLast)
		} else {
			assert.Equal(t, false, isLast)
		}
		assert.Equal(t, expected["id"], tableTuple["id"])
		assert.Equal(t, expected["name"], tableTuple["name"])
	}

	// Read again, result should be the same
	for i, expected := range tuples {
		tableTuple, isLast, err := mt.Read()
		assert.NoError(t, err)
		if i == len(tuples)-1 {
			assert.Equal(t, true, isLast)
		} else {
			assert.Equal(t, false, isLast)
		}
		assert.Equal(t, expected["id"], tableTuple["id"])
		assert.Equal(t, expected["name"], tableTuple["name"])
	}

	t.Run("Test IsEmpty", func(t *testing.T) {
		multiset1 := &Multiset{
			Name:  "multiset 1",
			Table: &table.MockTable{},
		}
		assert.Equal(t, true, multiset1.IsEmpty())

		multiset2 := &Multiset{
			Name:  "multiset 2",
			Table: mt,
		}
		assert.Equal(t, false, multiset2.IsEmpty())
	})

	t.Run("Test Fetch", func(t *testing.T) {
		multiset := &Multiset{
			Name:  "multiset",
			Table: mt,
		}
		dest := make(map[string]tuple.Value)
		for i, expected := range tuples {
			isLast, err := multiset.Fetch(dest)
			assert.NoError(t, err)
			if i == len(tuples)-1 {
				assert.Equal(t, true, isLast)
			} else {
				assert.Equal(t, false, isLast)
			}
			assert.Equal(t, expected["id"], dest["multiset.id"])
			assert.Equal(t, expected["name"], dest["multiset.name"])
		}
		// Fetch again to see it starts over
		for i, expected := range tuples {
			isLast, err := multiset.Fetch(dest)
			assert.NoError(t, err)
			if i == len(tuples)-1 {
				assert.Equal(t, true, isLast)
			} else {
				assert.Equal(t, false, isLast)
			}
			assert.Equal(t, expected["id"], dest["multiset.id"])
			assert.Equal(t, expected["name"], dest["multiset.name"])
		}
	})
}
