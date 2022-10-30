package tuple

import (
	"encoding/binary"
	"fmt"
)

type Value interface {
	ValueType() ValueType
	Serialize() []byte
}

type ValueType uint8

const (
	IntegerType ValueType = iota + 1
	StringType
)

type IntegerValue int64

func (i IntegerValue) ValueType() ValueType {
	return IntegerType
}

func (i IntegerValue) Serialize() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

type StringValue string

func (s StringValue) ValueType() ValueType {
	return StringType
}

func (s StringValue) Serialize() []byte {
	bareBytes := []byte(string(s))
	nBytesForSize := 4
	b := make([]byte, len(bareBytes)+nBytesForSize)
	binary.BigEndian.PutUint32(b, uint32(len(bareBytes)))
	copy(b[nBytesForSize:], bareBytes)
	return b
}

// returns the value deserialized from the byte slice, number of bytes read, and error
func DeserializeValue(b []byte, valueType ValueType) (Value, int, error) {
	switch valueType {
	case IntegerType:
		minLen := 8
		if len(b) < minLen {
			return nil, 0, fmt.Errorf("byte slice provided is not long enough for string type, min %d, received %d", minLen, len(b))
		}
		return IntegerValue(binary.BigEndian.Uint64(b)), minLen, nil
	case StringType:
		minLen := 4
		if len(b) < minLen {
			return nil, 0, fmt.Errorf("byte slice provided is not long enough for string type, min %d, received %d", minLen, len(b))
		}
		stringLen := int(binary.BigEndian.Uint32(b))
		if len(b) < minLen+stringLen {
			return nil, 0, fmt.Errorf("byte slice provided is not long enough for the size specified, min %d, received %d", minLen+stringLen, len(b))
		}
		return StringValue(string(b[minLen : minLen+stringLen])), stringLen + minLen, nil
	default:
		return nil, 0, fmt.Errorf("not a valid value type: %v", valueType)
	}
}
