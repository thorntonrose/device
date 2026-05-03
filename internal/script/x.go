package script

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

// X[n.c] -- copy source buffer to destination buffer (moving read pointer)

// n: characters to copy (default: 0 [all])
// c: stop character (default: 0 [end of buffer])
type X struct {
	Command
}

func NewX(memory *mem.Memory) X {
	return X{NewCommand(memory)}
}

func (self X) Run(parameters []string) int {
	n := self.NonNegative("n (characters)", parameters, 0, 0)
	c := self.Range("c (stop character)", parameters, 1, 0, 0, 255)

	return self.Copy(n, byte(c))
}

func (self X) Copy(n int, c byte) int {
	data := self.Memory.ReadAll(self.Memory.Source, n, byte(c))
	log.Printf("X.Copy: dest: %d, data: %v\n", self.Memory.Destination, data)
	self.Memory.WriteAll(self.Memory.Destination, data)

	return 0
}
