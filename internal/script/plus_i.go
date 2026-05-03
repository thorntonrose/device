package script

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// +I[a.t.s.n] -- send / receive data
//
// a: action (0 | 5, default: 0); 0 = send transmit buffer, 5 = wait for n characters then append to receive buffer
// t: reserved (default: 0)
// s: commands to skip if data received (default: 0)
// n: characters to wait for (default: 1)
type PlusI struct {
	Command
}

func NewPlusI(memory *mem.Memory) PlusI {
	return PlusI{NewCommand(memory)}
}

func (self PlusI) Run(parameters []string) (skip int) {
	log.Printf("PlusI.Run: %v\n", parameters)
	a := self.Code("a (action)", parameters, 0, 0, []int{0, 5})
	t := self.Int("t (reserved)", parameters, 1, 0)
	s := self.Int("s (skip)", parameters, 2, 0)
	n := self.Range("n (characters)", parameters, 3, 1, 1, mem.MaxBufferSize)

	return etc.If(a == 0, self.Transmit, func() int { return self.Receive(t, s, n) })()
}

func (self PlusI) Transmit() int {
	data := string(*self.Memory.Buffers[mem.Transmit])
	log.Printf("PlusI.Transmit: %v\n", data)
	fmt.Print(data)

	return 0
}

func (self *PlusI) Receive(t int, s int, n int) int {
	defer etc.Recover(func(e error) {})

	data := make([]byte, n)
	count := etc.Must(io.ReadFull(os.Stdin, data))
	log.Printf("PlusI.Receive: %v\n", data[:count])
	self.Memory.WriteAll(mem.Receive, data[:count])

	return s
}
