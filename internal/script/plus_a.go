package script

import (
	"log"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// +A[#v.a.c.s] -- compare variable
//
// v: variable (default: 0)
// a: comparison operation (default: 0); 0 = equal, 1 = greater than, 2 = less than
// c: constant value
// s: commands to skip if true
type PlusA struct {
	Command
}

func NewPlusA(memory *mem.Memory) PlusA {
	return PlusA{NewCommand(memory)}
}

func (self PlusA) Run(parameters []string) int {
	log.Printf("PlusA.Run: %v\n", parameters)
	v := self.Variable("v (variable)", parameters, 0, 0)
	a := self.Range("a (comparison)", parameters, 1, 0, 0, 2)
	c := self.Int("c (constant)", parameters, 2, 0)
	s := self.Int("s (skip)", parameters, 3, 0)

	return etc.If(self.Compare(self.Memory.Variables[v], a, c), s, 0)
}

func (self PlusA) Compare(v int, a int, c int) bool {
	return (a == 0 && v == c) || (a == 1 && v > c) || (a == 2 && v < c)
}
