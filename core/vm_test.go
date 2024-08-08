package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestVM(t *testing.T) {

	// 1 + 2 = 3
	// 1
	// push stack
	// 2
	// push stack
	// add
	// 3
	// push stack
	contractState := NewState()

	// data := []byte{0x02, 0x0a, 0x02, 0x0a, 0x0b}
	// result := vm.stack.Pop().(int)
	// assert.Equal(t, 4, result)

	// data := []byte{0x03, 0x0a, 0x02, 0x0a, 0x0e}
	// data := []byte{0x03, 0x0a, 0x46, 0x0c, 0x4f, 0x0c, 0x4f, 0x0c, 0x0d}
	// data := []byte{0x03, 0x0a, 0x61, 0x0c, 0x61, 0x0c, 0x61, 0x0c, 0x0d}


	// vm := NewVM(data)
	// assert.Nil(t, vm.Run())
	// fmt.Printf("stack: %v\n", vm.stack.data)
	// result := vm.stack.Pop().([]byte)
	// assert.Equal(t, []byte{0x61, 0x61, 0x61}, result)


	// data := []byte{0x03, 0x0a, 0x46, 0x0c, 0x4f, 0x0c, 0x4f, 0x0c, 0x0d}
	// vm := NewVM(data, contractState)
	// assert.Nil(t, vm.Run())
	// result := vm.stack.Pop().([]byte)
	// assert.Equal(t, "FOO", string(result))


	data := []byte{0x03, 0x0a, 0x2, 0xa, 0x0e}
	vm := NewVM(data, contractState)
	assert.Nil(t, vm.Run())
	result := vm.stack.Pop().(int)
	assert.Equal(t, 1, result)
}

func TestStack(t *testing.T) {
	s := NewStack(128)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	value := s.Pop()
	assert.Equal(t, 1, value)

	value = s.Pop()

	assert.Equal(t, 2, value)

	fmt.Printf("stack: %v\n", s.data)
}

// func TestStackStore(t *testing.T) {

// 	// F O O  => pack [F O O]
// 	data := []byte{0x03, 0x0a, 0x46, 0x0c, 0x4f, 0x0c, 0x4f, 0x0c, 0x0d, 0x05, 0x0a, 0x0f}
// 	// Push FOO to stack( key )
// 	// push 3 to the stack
// 	// Push 2 to the stack
// 	// 3 - 1
// 	// 1 is in the stack
// 	// [FOO, 1]
// 	// store 1 to the key FOO

// 	contractState := NewState()
// 	vm := NewVM(data, contractState)
// 	assert.Nil(t, vm.Run())

// 	fmt.Printf("stack: %v\n", vm.stack.data)
// 	fmt.Printf("state: %v\n", contractState.data) // [FOO:[5 0 0 0 0 0 0 0]]

// 	key, err := vm.contractState.Get([]byte("FOO"))
// 	assert.Nil(t, err)
// 	assert.Equal(t, key , []byte{5, 0, 0, 0, 0, 0, 0, 0})

// 	value := DeserializeInt64(key)
// 	assert.Equal(t, value, int64(5))
// }

func TestVM2(t *testing.T) {
	contractState := NewState()
	data := []byte{0x02, 0x0a, 0x03, 0x0a, 0x0b, 0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0x0f}
	otherData := []byte{0x02, 0x0a, 0x03, 0x0a, 0x0b, 0x4d, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0x0f}

	data = append(data, otherData...)

	vm := NewVM(data, contractState)
	assert.Nil(t, vm.Run())

	fmt.Printf("%+v\n", vm.stack.data)
	fmt.Printf("%+v\n", vm.contractState)

	valueBytes, err := vm.contractState.Get([]byte("FOO"))
	assert.Nil(t, err)
	value := DeserializeInt64(valueBytes)
	assert.Equal(t, value, int64(5))
}

// 2 + 3 = 5
// F O O  ->  key
// FOO = 5
func TestVM2Get(t *testing.T) {
	data := []byte{0x02, 0x0a, 0x03, 0x0a, 0x0b, 0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0x0f}
	pushFoo := []byte{0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0xae}

	data = append(data, pushFoo...)
	contractState := NewState()
	vm := NewVM(data, contractState)
	assert.Nil(t, vm.Run())

	fmt.Printf("%+v\n", vm.stack.data)

	value := vm.stack.Pop().([]byte)
	valueSerialized := DeserializeInt64(value)

	assert.Equal(t, int64(5), valueSerialized)
}


func TestMul(t *testing.T) {
	data := []byte{0x03, 0x0a, 0x02, 0x0a, 0xea}
	contractState := NewState()
	vm := NewVM(data, contractState)
	assert.Nil(t, vm.Run())

	fmt.Printf("%+v\n", vm.stack.data)

	value := vm.stack.Pop().(int)
	assert.Equal(t, 1, value)
}
