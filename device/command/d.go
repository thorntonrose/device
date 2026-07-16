package command

import (
	"log"

	"github.com/thorntonrose/device/device/mem"
)

// D[n] -- delete characters from end of destination buffer
type D struct {
	Command
}

func NewD(memory *mem.Memory) D {
	return D{New(memory)}
}

func (self D) Run(parameters []string) (skip int) {
	n := self.Positive("n (number of characters)", parameters, 0, 1)
	log.Printf("D: n: %d\n", n)

	self.Trim(mem.Transmit, n)
	return
}

func (self *D) Trim(bufNum int, n int) {
	self.Memory.Pointers[bufNum] = max(self.Memory.Pointers[bufNum]-n, 0)
	self.Memory.Slots[bufNum+1] = self.Memory.Slots[bufNum+1][:self.Memory.Pointers[bufNum]]
}
