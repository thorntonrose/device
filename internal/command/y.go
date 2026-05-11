package command

import (
	"fmt"
	"log"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
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
	log.Printf("Y.Run: %v\n", parameters)
	s := self.Int("s (skip)", parameters, 0, 0)

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
		log.Printf("Y.Append: slotNum: %d, data: %s\n", slotNum, string(data))
		self.Memory.WriteAll(self.Memory.Destination, self.Format(slotNum, data))
	}

	return slotNum
}

func (self *Y) Format(slotNum int, data []byte) []byte {
	return []byte(fmt.Sprintf("%03d=%s", slotNum, string(data)))
}
