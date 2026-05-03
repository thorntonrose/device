package script

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestA(t *testing.T) {
	a := NewA(mem.New())
	a.Memory.Set(0, []byte("FOO"))

	assert.Equal(t, 0, a.Run([]string{}))
	assert.Equal(t, []byte("FOO"), *a.Memory.Buffers[mem.Transmit])
}

func TestA_Slot(t *testing.T) {
	a := NewA(mem.New())
	a.Memory.Set(3, []byte("FOO"))

	assert.Equal(t, 0, a.Run([]string{"3"}))
	assert.Equal(t, []byte("FOO"), *a.Memory.Buffers[mem.Transmit])

	assert.Panics(t, func() { NewA(mem.New()).Run([]string{fmt.Sprintf("%d", mem.MaxSlots)}) })
}

func TestA_Skip(t *testing.T) {
	a := NewA(mem.New())
	a.Memory.Set(0, []byte("FOO"))
	a.Memory.Set(1, []byte(""))

	assert.Equal(t, 0, a.Run([]string{"", "1"}))  // not empty
	assert.Equal(t, 1, a.Run([]string{"1", "1"})) // empty
}
