package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

// B[b1.b2] -- Set source and destination buffers
type B struct {
	Command
}

func NewB(memory *mem.Memory) B {
	return B{New(memory)}
}

func (self B) Run(parameters []string) (skip int) {
	b1 := self.Code("b1 (source)", parameters, 0, 0, []int{0, 1, 2, 3, 4, 5, 9})
	b2 := self.Code("b2 (destination)", parameters, 1, 0, []int{0, 1, 2, 3, 4, 5})
	log.Printf("B.Run: b1: %d, b2: %d\n", b1, b2)

	self.SetSource(b1)
	self.SetDestination(b2)
	return
}

func (self B) SetSource(b int) {
	if b > 0 {
		self.SetSourceBuf(b)
		log.Printf("B.SetSource: reset: %d\n", self.Memory.Source)
		self.Memory.Reset(self.Memory.Source)
	}
}

func (self B) SetSourceBuf(b int) {
	if b != 9 {
		log.Printf("B.SetSource: %d\n", b)
		self.Memory.Source = b
	}
}

func (self B) SetDestination(b int) {
	if b > 0 {
		log.Printf("B.SetDestination: %d\n", b)
		self.Memory.Destination = b
	}
}
