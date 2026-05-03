package script

import (
	"log"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// A[m.s] -- append data to destination buffer
//
// m: memory slot (default: 0)
// s: commands to skip if slot is empty (default: 0)
type A struct {
	Command
}

func NewA(memory *mem.Memory) A {
	return A{NewCommand(memory)}
}

func (self A) Run(parameters []string) (skip int) {
	log.Printf("A.Run: %v\n", parameters)
	m := self.Range("m (memory slot)", parameters, 0, 0, 0, mem.MaxSlots-1)
	s := self.Int("s (skip)", parameters, 1, 0)

	return etc.If(len(self.Append(m)) == 0, s, 0)
}

func (self A) Append(m int) (data []byte) {
	if data = self.Memory.Slots[m]; len(data) > 0 {
		log.Printf("A.Append: dest: %d, data: %s\n", self.Memory.Destination, string(data))
		self.Memory.WriteAll(self.Memory.Destination, data)
	}

	return data
}
