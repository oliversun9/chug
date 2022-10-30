package btree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeyValueCell(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		key   []byte
		value []byte
	}{
		{
			name:  "both empty",
			key:   []byte{},
			value: []byte{},
		},
		{
			name:  "key empty",
			key:   []byte{},
			value: []byte{1, 2, 3, 4, 5, 6, 7, 1, 1, 1, 1, 1},
		},
		{
			name:  "value empty",
			key:   []byte{1, 2, 3, 4, 5, 6, 7},
			value: []byte{},
		},
		{
			name:  "both non-empty",
			key:   []byte{1, 2, 3, 4, 5},
			value: []byte{8, 4, 2, 1, 1},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			b := make([]byte, 1000)
			offset := 199
			kvc, err := newKeyValueCell(testCase.key, testCase.value)
			require.NoError(t, err)
			require.Equal(t, testCase.key, kvc.key())
			require.Equal(t, testCase.value, kvc.value())

			n, err := kvc.serialize(b[offset:])
			require.NoError(t, err)
			require.Equal(t, kvc.size(), n)

			newCell, err := deserializeCell(b[offset:])
			require.NoError(t, err)
			require.Equal(t, kvc.data, newCell.data)
		})
	}
}
