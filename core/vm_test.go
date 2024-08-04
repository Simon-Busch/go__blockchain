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

	data := []byte{0x02, 0x0a, 0x02, 0x0a, 0x0b}
	// data := []byte{0x03, 0x0a, 0x02, 0x0a, 0x0e}
	// data := []byte{0x03, 0x0a, 0x46, 0x0c, 0x4f, 0x0c, 0x4f, 0x0c, 0x0d}

	vm := NewVM(data)
	assert.Nil(t, vm.Run())
	fmt.Printf("stack: %v\n", vm.stack.data)
	result := vm.stack.Pop().(int)

	assert.Equal(t, 4, result)
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
