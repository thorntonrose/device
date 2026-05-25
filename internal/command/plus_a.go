package command

import (
	"log"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// +A[#v.a.c.s] -- compare variable
type PlusA struct {
	Command
}

func NewPlusA(memory *mem.Memory) PlusA {
	return PlusA{New(memory)}
}

func (self PlusA) Run(parameters []string) (skip int) {
	v := self.Variable("#v (variable)", parameters, 0, 0)
	a := self.Code("a (comparison)", parameters, 1, 0, []int{0, 1, 2})
	c := self.Int("c (constant)", parameters, 2, 0)
	s := self.Int("s (skip)", parameters, 3, 0)
	log.Printf("+A: v: %d, a: %d, c: %d, s: %d\n", v, a, c, s)

	return If(self.Compare(self.Memory.Variables[v], a, c), s, 0)
}

func (self PlusA) Compare(val int, a int, c int) (result bool) {
	result = (a == 0 && val == c) || (a == 1 && val > c) || (a == 2 && val < c)
	log.Printf("Compare: val: %d, a: %d, c: %d, result: %t\n", val, a, c, result)

	return
}
