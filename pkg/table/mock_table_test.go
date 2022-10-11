package table

import (
	"testing"

	"github.com/oliversun9/chug/pkg/algorithm"
	"github.com/oliversun9/chug/pkg/tuple"
	"github.com/stretchr/testify/assert"
)

func TestMockTable(t *testing.T) {
	t.Run("Test Read", func(t *testing.T) {
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
		mt, err := NewMockTable(
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
	})
}
