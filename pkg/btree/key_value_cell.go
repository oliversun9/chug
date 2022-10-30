package btree

import "github.com/oliversun9/chug/pkg/tuple"

type keyValueCell struct {
	data []byte
}

func newKeyValueCell(key tuple.Value, value []byte) keyValueCell {
	return keyValueCell{}
}

func (kvc keyValueCell) Key() tuple.Value {
	return nil
}

func (kvc keyValueCell) EncodedTuple() ([]byte, error) {
	return nil, nil
}
