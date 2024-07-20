package types

type Address [20]uint8

func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}
	return b
}

func NewAddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		panic("Given bytes should be 20")
	}

	var value [20]uint8
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}
	return Address(value)
}
