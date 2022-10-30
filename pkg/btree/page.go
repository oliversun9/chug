package btree

import "github.com/oliversun9/chug/pkg/tuple"

// a page stores a certain number of key value pairs
type Page interface {
	AddKeyValuePair([]byte, []byte) error
	GetValueByKey([]byte) ([]byte, error)
}

type pageID uint32

type cellPage4KB struct {
	data [4096]byte
}

type PageHeader struct {
}

func (p *cellPage4KB) header() PageHeader {
	return PageHeader{}
}

// TODO: still need to decide whether this is addTuple or addCell
func (p *cellPage4KB) addTuple(t tuple.Tuple) error {
	return nil
}

func (p *cellPage4KB) findTuple(key tuple.Value) (tuple.Tuple, error) {
	return nil, nil
}
