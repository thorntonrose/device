package command

import (
	"fmt"
	"log"

	. "github.com/thorntonrose/device/device/etc"
	"github.com/thorntonrose/device/device/mem"
)

// T[m.o.c.s] -- do math operation on contents of memory slot
type T struct {
	Command
}

func NewT(memory *mem.Memory) T {
	return T{New(memory)}
}

func (self T) Run(parameters []string) (skip int) {
	m := self.Code("m (memory slot)", parameters, 0, 0, []int{0, 20})
	o := self.Code("o (operation)", parameters, 1, 0, []int{0, 1, 2, 3, 4})
	c := self.Int("c (constant)", parameters, 2, 1)
	s := self.Int("s (skip)", parameters, 3, 0)
	log.Printf("T: m: %d, o: %d, c: %d, s: %d\n", m, o, c, s)

	return If(self.DoMath(m, o, c) == 0, s, 0)
}

func (self T) DoMath(m, o, c int) (val int) {
	val = self.Calculate(self.FromString(self.Memory.Slots[m]), o, c)
	log.Printf("DoMath: val: %d\n", val)
	self.Memory.Slots[m] = self.ToString(val)

	return
}

func (self T) FromString(data []byte) (result int) {
	fmt.Sscanf(string(data), "%d", &result)
	return
}

func (self T) ToString(val int) []byte {
	return []byte(fmt.Sprintf("%d", val))
}
