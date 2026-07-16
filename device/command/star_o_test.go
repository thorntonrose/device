package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/device/mem"
	_ "github.com/thorntonrose/device/device/testing"
)

func TestStarO(t *testing.T) {
	so := NewStarO(mem.New())

	assert.Equal(t, 0, so.Run([]string{}))
	assert.Equal(t, 1, so.Memory.Variables[0])

	assert.Panics(t, func() { so.Run([]string{"", "5"}) })
}

func TestStarO_Subtract(t *testing.T) {
	so := NewStarO(mem.New())
	so.Memory.Variables[0] = 1

	assert.Equal(t, 0, so.Run([]string{"#0", "1"}))
	assert.Equal(t, 0, so.Memory.Variables[0])
}

func TestStarO_Multiply(t *testing.T) {
	so := NewStarO(mem.New())
	so.Memory.Variables[0] = 1

	assert.Equal(t, 0, so.Run([]string{"", "2", "2"}))
	assert.Equal(t, 2, so.Memory.Variables[0])
}

func TestStarO_Divide(t *testing.T) {
	so := NewStarO(mem.New())
	so.Memory.Variables[0] = 2

	assert.Equal(t, 0, so.Run([]string{"", "3", "2"}))
	assert.Equal(t, 1, so.Memory.Variables[0])
}

func TestStarO_Modulo(t *testing.T) {
	so := NewStarO(mem.New())
	so.Memory.Variables[0] = 5

	assert.Equal(t, 0, so.Run([]string{"", "4", "2"}))
	assert.Equal(t, 1, so.Memory.Variables[0])
}

func TestStarO_Skip(t *testing.T) {
	so := NewStarO(mem.New())
	so.Memory.Variables[0] = 1

	assert.Equal(t, 1, so.Run([]string{"", "1", "", "1"}))
}
