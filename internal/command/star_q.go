package command

import (
	"fmt"
	"log"
	"math/big"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// Q[#v.b.a] -– set variable from buffer
type StarQ struct {
	Command
}

func NewStarQ(memory *mem.Memory) StarQ {
	return StarQ{Command: New(memory)}
}

func (self StarQ) Run(parameters []string) (skip int) {
	v := self.Variable("#v (variable)", parameters, 0, 0)
	b := self.Code("b (buffer)", parameters, 1, 0, []int{0, 1})
	a := self.Code("a (conversion)", parameters, 2, 0, []int{0, 1})
	log.Printf("*Q: v: %d, b: %d, a: %d\n", v, b, a)

	data := self.Memory.Slots[If(b == 0, self.Memory.Destination, b)+1]
	self.Memory.Variables[v] = If(a == 0, self.FromString(data), self.FromASCII(data))
	log.Printf("*Q: #%d: %d\n", v, self.Memory.Variables[v])

	return
}

func (self StarQ) FromString(data []byte) (result int) {
	fmt.Sscanf(string(data), "%d", &result)
	return
}

func (self StarQ) FromASCII(data []byte) int {
	return int(big.NewInt(0).SetBytes(data).Int64())
}
