package command

import (
	"fmt"
	"log"

	. "github.com/thorntonrose/device/device/etc"
	"github.com/thorntonrose/device/device/mem"
	"github.com/thorntonrose/device/device/mem/buf"
)

// Y[s] -- append next non-empty memory slot to destination buffer
type Y struct {
	Command
	SlotNum int
}

func NewY(memory *mem.Memory) *Y {
	return &Y{New(memory), 0}
}

func (self *Y) Run(parameters []string) (skip int) {
	s := self.Int("s (skip)", parameters, 0, 0)
	log.Printf("Y: s: %d\n", s)

	return If(self.AppendNextNonEmpty() == 0, s, 0)
}

func (self *Y) AppendNextNonEmpty() int {
	self.SlotNum = self.Append(self.NextNonEmpty()) + 1
	return self.SlotNum
}

func (self *Y) NextNonEmpty() (int, []byte) {
	for i := self.SlotNum; i < len(self.Memory.Slots); i++ {
		if data := self.Memory.Slots[i]; len(data) > 0 {
			return i, data
		}
	}

	return -1, nil
}

func (self *Y) Append(slotNum int, data []byte) int {
	if slotNum != -1 {
		log.Printf("Append: slotNum: %d, data: %s\n", slotNum, string(data))
		buf.WriteAll(self.Memory, self.Memory.Destination, self.Format(slotNum, data))
	}

	return slotNum
}

func (self *Y) Format(slotNum int, data []byte) []byte {
	return []byte(fmt.Sprintf("%03d=%s", slotNum, string(data)))
}
