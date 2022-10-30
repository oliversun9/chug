package tuple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value Value
	}{
		{
			name:  "integer positive",
			value: IntegerValue(141413),
		},
		{
			name:  "integer 0",
			value: IntegerValue(0),
		},
		{
			name:  "integer negative",
			value: IntegerValue(-12314),
		},
		{
			name:  "string empty",
			value: StringValue(""),
		},
		{
			name:  "string short",
			value: StringValue("a"),
		},
		{
			name:  "string longer",
			value: StringValue("fjakljfjekwajfewlakfjawl1231414fjakwlejf90-9;-09-=[]';'l;'"),
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			b := testCase.value.Serialize()
			v, n, err := DeserializeValue(b, testCase.value.ValueType())
			require.NoError(t, err)
			require.Equal(t, testCase.value, v)
			require.Equal(t, len(b), n)
		})
	}
}
