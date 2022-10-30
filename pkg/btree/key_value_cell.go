package btree

import (
	"encoding/binary"
	"fmt"
)

const sizeLen = 4

type keyValueCell struct {
	/*
		|        4 |          4 | key_size | value_size |
		| key_size | value_size |      key |      value |
	*/
	data []byte
}

func newKeyValueCell(key []byte, value []byte) (keyValueCell, error) {
	data := make([]byte, sizeLen*2+len(key)+len(value))
	binary.BigEndian.PutUint32(data, uint32(len(key)))
	binary.BigEndian.PutUint32(data[sizeLen:], uint32(len(value)))
	offset := sizeLen * 2
	copied := copy(data[offset:], key)
	if copied != len(key) {
		return keyValueCell{}, fmt.Errorf("wrong number of bytes copied from key to cell, %d bytes expected, %d copied", len(key), copied)
	}
	offset += copied
	copied = copy(data[offset:], value)
	if copied != len(value) {
		return keyValueCell{}, fmt.Errorf("wrong number of bytes copied from value to cell, %d bytes expected, %d copied", len(value), copied)
	}
	return keyValueCell{
		data: data,
	}, nil
}

// basically reads up to where it ends
func deserializeCell(src []byte) (keyValueCell, error) {
	minLen := sizeLen * 2
	if len(src) < minLen {
		return keyValueCell{}, fmt.Errorf("src too short: at least %d bytes expected, %d bytes provided", minLen, len(src))
	}
	keyLen := int(binary.BigEndian.Uint32(src))
	valLen := int(binary.BigEndian.Uint32(src[sizeLen:]))
	totalLen := keyLen + valLen + minLen
	newCell := keyValueCell{
		make([]byte, totalLen),
	}
	copied := copy(newCell.data, src[:totalLen])
	if copied != totalLen {
		return newCell, fmt.Errorf("wrong number of bytes copied, %d expected, %d copied", totalLen, copied)
	}
	return newCell, nil
}

// called when defragmentation is needed, or when a new cell is created
func (kvc keyValueCell) serialize(dst []byte) (int, error) {
	if len(dst) < kvc.size() {
		return 0, fmt.Errorf("unable to serialize cell, require space of at least %d bytes, %d provided", kvc.size(), len(dst))
	}
	copied := copy(dst, kvc.data)
	if copied != kvc.size() {
		return copied, fmt.Errorf("wrong number of bytes copied: %d bytes expected, %d copied", kvc.size(), copied)
	}
	return copied, nil
}

func (kvc keyValueCell) size() int {
	return len(kvc.data)
}

func (kvc keyValueCell) key() []byte {
	keySize := int(binary.BigEndian.Uint32(kvc.data))
	offset := sizeLen * 2
	return kvc.data[offset : offset+keySize]
}

func (kvc keyValueCell) value() []byte {
	keySize := int(binary.BigEndian.Uint32(kvc.data))
	valueSize := int(binary.BigEndian.Uint32(kvc.data[sizeLen:]))
	offset := sizeLen*2 + keySize
	return kvc.data[offset : offset+valueSize]
}
