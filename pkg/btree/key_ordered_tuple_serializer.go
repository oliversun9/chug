package btree

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"

	"github.com/oliversun9/chug/pkg/tuple"
)

type keyOrderedTupleSerializer struct {
	schema      tuple.KeyedSchema
	orderedKeys []string
}

func newKeyOrderedTupleSerialzer(schema tuple.KeyedSchema) keyOrderedTupleSerializer {
	tupleSerializer := keyOrderedTupleSerializer{
		schema:      schema,
		orderedKeys: make([]string, len(schema.Schema())),
	}
	i := 0
	for name := range schema.Schema() {
		tupleSerializer.orderedKeys[i] = name
		i++
	}
	sort.Strings(tupleSerializer.orderedKeys)
	return tupleSerializer
}

func (s keyOrderedTupleSerializer) Serialize(t tuple.Tuple) ([]byte, error) {
	if err := s.schema.Schema().ValidateTuple(t); err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	// at the beginning, store the size of the whole byte slice (including this siz)
	byteCountSize := 4
	n, err := buf.Write(make([]byte, byteCountSize))
	if err != nil {
		return nil, err
	}
	if n != byteCountSize {
		return nil, fmt.Errorf("error writing %d bytes into buffer, wrote %d", byteCountSize, n)
	}
	for _, keyName := range s.orderedKeys {
		if _, err = buf.Write(t[keyName].Serialize()); err != nil {
			return nil, err
		}
	}
	b := buf.Bytes()
	if len(b) < byteCountSize {
		return nil, fmt.Errorf("buffer yields byte slice of size %d, expecting at least %d", len(b), byteCountSize)
	}
	binary.BigEndian.PutUint32(b, uint32(len(b)))
	return b, nil
}

func (s keyOrderedTupleSerializer) Deserialize(b []byte, t tuple.Tuple) error {
	offset := 4
	if len(b) < offset {
		return fmt.Errorf("byte slices has length %d less than minimal, %d", len(b), offset)
	}
	byteLength := binary.BigEndian.Uint32(b)
	if len(b) < int(byteLength) {
		return fmt.Errorf("indicated %d bytes, but the slice has length %d", byteLength, len(b))
	}
	for _, fieldName := range s.orderedKeys {
		if offset >= len(b) {
			return fmt.Errorf("attempting to read from index %d of byte slice with len %d", offset, len(b))
		}
		v, nBytesRead, err := tuple.DeserializeValue(b[offset:], s.schema.Schema()[fieldName])
		if err != nil {
			return err
		}
		offset += nBytesRead
		t[fieldName] = v
	}
	if err := s.schema.Schema().ValidateTuple(t); err != nil {
		return err
	}
	return nil
}
