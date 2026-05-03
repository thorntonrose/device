package script

import (
	"log"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// *O[#v.o.c.s] -- perform arithmetic on variable
//
// v: variable (default: 0)
// o: operation (0 - 4, default: 0); 0 = add, 1 = subtract, 2 = multiply, 3 = divide, 4 = modulo
// c: constant value
// s: commands to skip if result is 0
type StarO struct {
	Command
}

func NewStarO(memory *mem.Memory) StarO {
	return StarO{NewCommand(memory)}
}

func (self StarO) Run(parameters []string) (skip int) {
	log.Printf("StarO.Run: %v\n", parameters)
	v := self.Variable("v (variable)", parameters, 0, 0)
	o := self.Range("o (operation)", parameters, 1, 0, 0, 4)
	c := self.Int("c (constant)", parameters, 2, 1)
	s := self.Int("s (skip)", parameters, 3, 0)

	return self.PerformArithmetic(v, o, c, s)
}

func (self StarO) PerformArithmetic(v, o, c, s int) int {
	result := self.Calculate(self.Memory.Variables[v], o, c)
	log.Printf("StarO.PerformArithmetic: result: %d\n", result)
	self.Memory.Variables[v] = result

	return etc.If(result == 0, s, 0)
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
