package command

import (
	"fmt"
	"io"
	"log"
	"os"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// +I[a.t.s.n] -- send / receive data
type PlusI struct {
	Command
}

func NewPlusI(memory *mem.Memory) PlusI {
	return PlusI{New(memory)}
}

func (self PlusI) Run(parameters []string) (skip int) {
	log.Printf("PlusI.Run: %v\n", parameters)
	a := self.Code("a (action)", parameters, 0, 0, []int{0, 1, 5})
	t := self.Int("t (reserved)", parameters, 1, 0)
	s := self.Int("s (skip)", parameters, 2, 0)
	n := self.Range("n (characters)", parameters, 3, 1, 1, mem.MaxBufferSize)

	return If(a == 5, func() int { return self.Receive(t, s, n) }, func() int { return self.Transmit(a) })()
}

func (self PlusI) Transmit(a int) (skip int) {
	data := string(*self.Memory.Buffers[mem.Transmit])
	log.Printf("PlusI.Transmit: %v\n", data)
	fmt.Print(data + If(a == 1, "\n", ""))

	return
}

func (self *PlusI) Receive(t int, s int, n int) (skip int) {
	defer Recover(func(e error) {})

	data := make([]byte, n)
	count := Must(io.ReadFull(os.Stdin, data))
	log.Printf("PlusI.Receive: %v\n", data[:count])
	self.Memory.WriteAll(mem.Receive, data[:count])

	return s
}
