package script

import "github.com/thorntonrose/device/internal/mem"

type StarN struct {
	Command
}

// *N -- set variable to constant value
//
// #v -- variable to set (default: 0)
// c -- constant value (default: 0)
func NewStarN(memory *mem.Memory) StarN {
	return StarN{NewCommand(memory)}
}

func (self StarN) Run(parameters []string) int {
	v := self.Variable("c ()", parameters, 0, 0)
	c := self.Int("", parameters, 1, 0)

	self.Memory.Variables[v] = c
	return 0
}
