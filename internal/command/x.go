package command

import "github.com/thorntonrose/device/internal/mem"

// X<n>.<c> -- copy source buffer to destination buffer (moving read pointer)
// n: number of characters to copy (default: 0 [all])
// c: stop character (default: 0 [end of buffer])
type X struct {
	Command
}

func NewX(memory *mem.Memory) X {
	return X{New(memory)}
}

func (x X) Run(parameters []string) int {
	n := x.Int("n (number to copy)", parameters, 0, 0)
	c := x.Int("c (stop character)", parameters, 1, 0)
	x.Memory.WriteAll(x.Memory.Destination, x.Memory.ReadAll(x.Memory.Source, n, byte(c)))

	return 0
}
