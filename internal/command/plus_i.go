package command

import (
	"fmt"
	"io"
	"log"
	"os"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
	"github.com/thorntonrose/device/internal/mem/buf"
)

// +I[a.t.s.n] -- send / receive data
type PlusI struct {
	Command
}

func NewPlusI(memory *mem.Memory) PlusI {
	return PlusI{New(memory)}
}

func (self PlusI) Run(parameters []string) (skip int) {
	a := self.Code("a (action)", parameters, 0, 0, []int{0, 1, 5})
	t := self.Int("t (reserved)", parameters, 1, 0)
	s := self.Int("s (skip)", parameters, 2, 0)
	n := self.Range("n (characters)", parameters, 3, 1, 1, mem.MaxBufferSize)
	log.Printf("+I: a: %d, t: %d, s: %d, n: %d\n", a, t, s, n)

	return If(a == 5, func() int { return self.Receive(t, s, n) }, func() int { return self.Transmit(a) })()
}

func (self PlusI) Transmit(a int) (skip int) {
	data := string(*self.Memory.Buffers[mem.Transmit])
	log.Printf("Transmit: %v (%s)\n", []byte(data), data)
	fmt.Print(data + If(a == 1, "\n", ""))

	return
}

func (self *PlusI) Receive(t int, s int, n int) (skip int) {
	data := Must(io.ReadAll(io.LimitReader(os.Stdin, int64(n))))
	log.Printf("Receive: %v (%s)\n", data, string(data))

	if len(data) > 0 {
		buf.WriteAll(self.Memory, mem.Receive, data)
		return s
	}

	return 0
}
