package command

import (
	"testing"

	"github.com/thorntonrose/device/internal/mem"
)

func TestPlusI_RunAction(t *testing.T) {
	c := NewPlusI(mem.New())
	*c.Memory.Buffers[mem.Transmit] = []byte("HELLO")

	c.Run([]string{})
	// c.RunAction(5, 0, 0, 0)
}
