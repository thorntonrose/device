package vars

import (
	"fmt"
	"log"
	"math/big"

	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// +Q[#v.a] – append variable to destination buffer
//
// #v: variable (default: #0)
// a: conversion code (0 - 1, default: 0); 0 = string (2748 -> “2748”), 1 = ASCII (86 -> “V”, 22617 -> “XY”)
type PlusQ struct {
	command.Command
}

func NewPlusQ(memory *mem.Memory) PlusQ {
	return PlusQ{Command: command.New(memory)}
}

func (self PlusQ) Run(parameters []string) int {
	log.Printf("PlusQ.Run: %v\n", parameters)
	v := self.Variable("v (variable)", parameters, 0, 0)
	a := self.Code("a (conversion)", parameters, 1, 0, []int{0, 1})

	self.Memory.WriteAll(self.Memory.Destination, etc.If(a == 0, self.String(v), self.ASCII(v)))
	return 0
}

func (self PlusQ) String(v int) []byte {
	return []byte(fmt.Sprintf("%d", self.Memory.Variables[v]))
}

func (self PlusQ) ASCII(v int) []byte {
	return big.NewInt(int64(self.Memory.Variables[v])).Bytes()
}
