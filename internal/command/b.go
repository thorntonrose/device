package command

import (
	"fmt"

	"github.com/thorntonrose/device/internal/mem"
)

// B[<s>][.<d>] -- Set source and destination buffers
// s: source (0 - <max-buffers> | 9, default: 0)
// d: destination (0 - <max-buffers> | 9, default: 0)
//
// 0 = no change
// 1 - <max-buffers> = set buffer; reset pointer
// 9 = reset pointer
type B struct {
	Command
}

func NewB(memory mem.Memory) B {
	return B{New(memory)}
}

func (b B) Run(parameters []string) int {
	b.Set(&b.Memory.Source, b.Buffer("source (s)", parameters, 0))
	b.Set(&b.Memory.Destination, b.Buffer("destination (d)", parameters, 1))

	return 0
}

func (b B) Set(curr *int, buf int) {
	if buf > 0 && buf < 9 {
		*curr = buf
	}

	b.Reset(*curr, buf)
}

func (b B) Reset(curr int, buf int) {
	if buf == 9 {
		b.Memory.Reset(curr)
	}
}

func (b B) Buffer(name string, parameters []string, index int) int {
	buf := b.ToInt(name, parameters, index, 0)

	if (buf >= 0 && buf <= mem.MaxBuffers) || buf == 9 {
		return buf
	}

	panic(fmt.Sprintf("%s: %d", name, buf))
}
