package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

// *N -- set variable to constant value
type StarN struct {
	Command
}

func NewStarN(memory *mem.Memory) StarN {
	return StarN{New(memory)}
}

func (self StarN) Run(parameters []string) (skip int) {
	v := self.Variable("#v (variable)", parameters, 0, 0)
	c := self.Int("c (constant)", parameters, 1, 0)
	log.Printf("*N: v: %d, c: %d\n", v, c)

	self.Memory.Variables[v] = c
	return
}
