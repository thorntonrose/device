package command

import (
	"log"

	. "github.com/thorntonrose/device/device/etc"
	"github.com/thorntonrose/device/device/mem"
)

// *O[#v.o.c.s] -- do math operation on variable
type StarO struct {
	Command
}

func NewStarO(memory *mem.Memory) StarO {
	return StarO{New(memory)}
}

func (self StarO) Run(parameters []string) (skip int) {
	v := self.Variable("#v (variable)", parameters, 0, 0)
	o := self.Code("o (operation)", parameters, 1, 0, []int{0, 1, 2, 3, 4})
	c := self.Int("c (constant)", parameters, 2, 1)
	s := self.Int("s (skip)", parameters, 3, 0)
	log.Printf("*O: v: %d, o: %d, c: %d, s: %d\n", v, o, c, s)

	return If(self.DoMath(v, o, c) == 0, s, 0)
}

func (self StarO) DoMath(v, o, c int) (val int) {
	val = self.Calculate(self.Memory.Variables[v], o, c)
	log.Printf("*DoMath: val: %d\n", val)
	self.Memory.Variables[v] = val

	return
}
