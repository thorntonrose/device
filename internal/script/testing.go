package script

import (
	"strconv"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

func NewTestScript(memory *mem.Memory) Script {
	script := NewScript(memory)
	script.Runners = map[string]Runner{"Z": NewZ(memory), "+Z": NewZ(memory), "*Z": NewZ(memory)}

	return script
}

type Z struct {
	Command
}

func NewZ(memory *mem.Memory) Z {
	return Z{NewCommand(memory)}
}

func (self Z) Run(parameters []string) int {
	if len(parameters) > 1 {
		*self.Memory.Buffers[mem.Transmit] = []byte(parameters[1])
	}

	return etc.Must(strconv.Atoi(parameters[0]))
}
