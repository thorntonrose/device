package bufs

import (
	"fmt"
	"log"

	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// Y[s] -- append next non-empty memory slot to destination buffer
//
// s: commands to skip after all memory slots are read
type Y struct {
	command.Command
	SlotNum int
}

func NewY(memory *mem.Memory) *Y {
	return &Y{command.New(memory), 0}
}

func (self *Y) Run(parameters []string) (skip int) {
	log.Printf("Y.Run: %v\n", parameters)
	s := self.Int("s (skip)", parameters, 0, 0)

	return self.AppendNextNonEmpty(s)
}

func (self *Y) AppendNextNonEmpty(s int) int {
	slotNum, data := self.Next()
	self.Append(slotNum, data)
	self.SlotNum = slotNum + 1

	return etc.If(self.SlotNum == 0, s, 0)
}

func (self *Y) Next() (int, []byte) {
	for i := self.SlotNum; i < len(self.Memory.Slots); i++ {
		if data := self.Memory.Slots[i]; len(data) > 0 {
			return i, data
		}
	}

	return -1, nil
}

func (self *Y) Append(slotNum int, data []byte) {
	if len(data) > 0 {
		log.Printf("Y.Append: slotNum: %d, data: %s\n", slotNum, string(data))
		self.Memory.WriteAll(self.Memory.Destination, self.Format(slotNum, data))
	}
}

func (self *Y) Format(slotNum int, data []byte) []byte {
	return []byte(fmt.Sprintf("%03d=%s", slotNum, string(data)))
}
