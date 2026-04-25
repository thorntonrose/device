package command

import (
	"fmt"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// +I[a.t.s.n] -- send / receive data
//
// a: action (0 | 5, default 0); 0 = send transmit buffer, 5 = wait for n characters then append to receive buffer
// t: receive timeout in seconds
// s: number of commands to skip it data received (default 0)
// n: number of characters to wait for (default 1)
type PlusI struct {
	Command
}

func NewPlusI(memory *mem.Memory) PlusI {
	return PlusI{New(memory)}
}

func (self PlusI) Run(parameters []string) (skip int) {
	a := self.Code("a (action)", parameters, 0, 0, []int{0, 5})
	t := self.NonNegative("t (timeout)", parameters, 1, 0)
	s := self.Int("s (skip)", parameters, 2, 0)
	n := self.Range("n (number of characters)", parameters, 3, 1, 1, mem.MaxBufferSize)

	return etc.If(a == 0, func() int { return self.Transmit() }, func() int { return self.Receive(t, s, n) })()
}

func (self PlusI) Transmit() int {
	fmt.Println(string(*self.Memory.Buffers[mem.Transmit]))
	return 0
}

func (self *PlusI) Receive(timeout int, skip int, n int) int {
	panic("not yet implemented")
}
