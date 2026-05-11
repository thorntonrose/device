package command

import (
	"log"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// *O[#v.o.c.s] -- perform arithmetic on variable
type StarO struct {
	Command
}

func NewStarO(memory *mem.Memory) StarO {
	return StarO{New(memory)}
}

func (self StarO) Run(parameters []string) (skip int) {
	log.Printf("StarO.Run: %v\n", parameters)
	v := self.Variable("#v (variable)", parameters, 0, 0)
	o := self.Code("o (operation)", parameters, 1, 0, []int{0, 1, 2, 3, 4})
	c := self.Int("c (constant)", parameters, 2, 1)
	s := self.Int("s (skip)", parameters, 3, 0)

	return If(self.PerformArithmetic(v, o, c) == 0, s, 0)
}

func (self StarO) PerformArithmetic(v, o, c int) (result int) {
	result = self.Calculate(self.Memory.Variables[v], o, c)
	log.Printf("StarO.PerformArithmetic: result: %d\n", result)
	self.Memory.Variables[v] = result

	return
}

func (self StarO) Calculate(v, o, c int) int {
	switch o {
	case 0:
		return v + c
	case 1:
		return v - c
	case 2:
		return v * c
	case 3:
		return v / c
	case 4:
		return v % c
	default:
		return 0
	}
}
