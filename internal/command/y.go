package command

import (
	"fmt"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// Y[s] -- append next non-empty memory slot to destination buffer
//
// s: commands to skip after all memory slots are read
type Y struct {
	Command
	SlotNum int
}

func NewY(memory *mem.Memory) *Y {
	return &Y{New(memory), 0}
}

func (c *Y) Run(parameters []string) (skip int) {
	s := c.Int("s (skip)", parameters, 0, 0)
	c.Append(c.SlotNum)
	c.Next()

	return etc.If(c.SlotNum == 0, s, 0)
}

func (c *Y) Append(slotNum int) {
	if data := c.Memory.Slots[slotNum]; len(data) > 0 {
		c.Memory.WriteAll(c.Memory.Destination, c.Format(slotNum, data))
	}
}

func (c *Y) Format(slotNum int, data []byte) []byte {
	return []byte(fmt.Sprintf("%03d:%s", slotNum, string(data)))
}

func (c *Y) Next() {
	c.SlotNum++

	if c.SlotNum == len(c.Memory.Slots) {
		c.SlotNum = 0
	}
}
