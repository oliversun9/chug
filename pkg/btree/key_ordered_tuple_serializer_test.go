package btree

import (
	"testing"

	"github.com/oliversun9/chug/pkg/tuple"
	"github.com/stretchr/testify/require"
)

func TestKeyOrderedTupleSerializer(t *testing.T) {
	t.Parallel()
	schema1, err := tuple.NewKeyedSchema(
		"name123",
		tuple.StringType,
		map[string]tuple.ValueType{
			"name123":       tuple.StringType,
			"rank":          tuple.IntegerType,
			"whateverIndex": tuple.IntegerType,
			"about":         tuple.StringType,
			"origin":        tuple.StringType,
			"someText":      tuple.StringType,
		},
	)
	require.NoError(t, err)
	tests := []struct {
		name   string
		schema tuple.KeyedSchema
		tuple  tuple.Tuple
	}{
		{
			name:   "any tuple",
			schema: schema1,
			tuple: tuple.Tuple{
				"name123":       tuple.StringValue("ok"),
				"rank":          tuple.IntegerValue(3123),
				"whateverIndex": tuple.IntegerValue(-141243),
				"about":         tuple.StringValue(""),
				"origin":        tuple.StringValue("somewhere"),
				"someText":      tuple.StringValue("yea, just some text, 1321r 4124 3123 4312"),
			},
		},
		{
			name:   "zero int",
			schema: schema1,
			tuple: tuple.Tuple{
				"name123":       tuple.StringValue("ok"),
				"rank":          tuple.IntegerValue(0),
				"whateverIndex": tuple.IntegerValue(0),
				"about":         tuple.StringValue(""),
				"origin":        tuple.StringValue("somewhere"),
				"someText":      tuple.StringValue("yea, just some text, 1321r 4124 3123 4312"),
			},
		},
		{
			name:   "empty string",
			schema: schema1,
			tuple: tuple.Tuple{
				"name123":       tuple.StringValue(""),
				"rank":          tuple.IntegerValue(0),
				"whateverIndex": tuple.IntegerValue(0),
				"about":         tuple.StringValue(""),
				"origin":        tuple.StringValue(""),
				"someText":      tuple.StringValue(""),
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			serializer := newKeyOrderedTupleSerialzer(testCase.schema)
			b, err := serializer.Serialize(testCase.tuple)
			require.NoError(t, err)
			dst := tuple.Tuple{}
			err = serializer.Deserialize(b, dst)
			require.NoError(t, err)
			require.Equal(t, testCase.tuple, dst)
		})
	}
}
