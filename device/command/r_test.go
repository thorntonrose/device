package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/device/mem"
	"github.com/thorntonrose/device/device/mem/buf"
)

func TestR(t *testing.T) {
	r := NewR(mem.New())

	AssertR(t, r, []string{""}, []byte{'\t'})
	AssertR(t, r, []string{"32"}, []byte{' '})
	AssertR(t, r, []string{"'A'"}, []byte("A"))
}

func AssertR(t *testing.T, r R, parameters []string, expect []byte) {
	buf.Clear(r.Memory, mem.Transmit)
	r.Run(parameters)
	assert.Equal(t, expect, r.Memory.Slots[mem.Transmit+1])
}
