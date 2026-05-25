package command

import (
	"fmt"
	"log"
	"math/big"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
	"github.com/thorntonrose/device/internal/mem/buf"
)

// +Q[#v.a] – append variable to destination buffer
type PlusQ struct {
	Command
}

func NewPlusQ(memory *mem.Memory) PlusQ {
	return PlusQ{Command: New(memory)}
}

func (self PlusQ) Run(parameters []string) (skip int) {
	v := self.Variable("#v (variable)", parameters, 0, 0)
	a := self.Code("a (conversion)", parameters, 1, 0, []int{0, 1})
	log.Printf("+Q: v: %d, a: %d\n", v, a)

	buf.WriteAll(self.Memory, self.Memory.Destination, If(a == 0, self.ToString(v), self.ToASCII(v)))
	return
}

func (self PlusQ) ToString(v int) []byte {
	return []byte(fmt.Sprintf("%d", self.Memory.Variables[v]))
}

func (self PlusQ) ToASCII(v int) []byte {
	return big.NewInt(int64(self.Memory.Variables[v])).Bytes()
}
