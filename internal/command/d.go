package command

import (
	"log"

	"github.com/thorntonrose/device/internal/mem"
)

// D[n] -- delete characters from end of destination buffer
type D struct {
	Command
}

func NewD(memory *mem.Memory) D {
	return D{New(memory)}
}

func (d D) Run(parameters []string) (skip int) {
	n := d.Positive("n (number of characters)", parameters, 0, 1)
	log.Printf("D.Run: n: %d\n", n)

	d.Memory.Trim(mem.Transmit, n)
	return
}
