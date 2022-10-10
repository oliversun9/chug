package tuple

type Value interface {
	valueType() ValueType
}

type ValueType uint8

const (
	IntegerType ValueType = iota + 1
	StringType
)

type IntegerValue int64

func (i IntegerValue) valueType() ValueType {
	return IntegerType
}

type StringValue string

func (s StringValue) valueType() ValueType {
	return StringType
}
