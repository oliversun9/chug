package btree

type KeyCell struct {
	data []byte
}

func (kc KeyCell) Serialize() []byte {
	return kc.data
}
