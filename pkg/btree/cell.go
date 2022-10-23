package btree

import "github.com/oliversun9/chug/pkg/tuple"

type Cell interface {
	Key() tuple.Value
}

type KeyPointerCell interface {
	Cell
	PageID() pageID
}

type KeyValueCell interface {
	Cell
	// Does not return a tuple, it shouldn't know how a tuple
	// is encoded, but only needs to know where to put the
	// encoded tuple within itself (the cell)
	EncodedTuple() []byte
}
