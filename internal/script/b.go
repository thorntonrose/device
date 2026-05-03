package script

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

// B[b1.b2] -- Set source and destination buffers
//
// b1: source buffer (0 - <max-buffers> | 9, default: 0); 0 = no change, 1 - <max-buffers> = set buffer and reset
// pointer; 9 = reset pointer
//
// b2: destination buffer (0 - <max-buffers>, default: 0); 0 = no change, 1 - <max-buffers> = set buffer
type B struct {
	Command
}

func NewB(memory *mem.Memory) B {
	return B{NewCommand(memory)}
}

func (self B) Run(parameters []string) int {
	log.Printf("B.Run: %v\n", parameters)
	b1 := self.Code("b1 (source)", parameters, 0, 0, []int{0, 1, 2, 3, 4, 5, 9})
	b2 := self.Code("b2 (destination)", parameters, 1, 0, []int{0, 1, 2, 3, 4, 5})

	return self.Set(b1, b2)
}

func (self B) Set(b1 int, b2 int) int {
	self.SetSource(b1)
	self.SetDestination(b2)

	return 0
}

func (self B) SetSource(b int) {
	if b > 0 && b != 9 {
		log.Printf("B.SetSource: %d\n", b)
		self.Memory.Source = b
	}

	log.Printf("B.SetSource: reset: %d\n", self.Memory.Source)
	self.Memory.Reset(self.Memory.Source)
}

func (self B) SetDestination(b int) {
	if b > 0 {
		log.Printf("B.SetDestination: %d\n", b)
		self.Memory.Destination = b
	}
}
