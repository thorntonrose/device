package script

import (
	"strconv"

	"github.com/thorntonrose/device/internal/command"
	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

func NewTestScript(memory *mem.Memory) Script {
	script := New(memory)
	script.Commands = map[string]Command{"Z": NewZ(memory), "+Z": NewZ(memory), "*Z": NewZ(memory)}

	return script
}

//-----------------------------------------------------------------------------

type Z struct {
	command.Command
}

func NewZ(memory *mem.Memory) Z {
	return Z{command.New(memory)}
}

func (self Z) Run(parameters []string) int {
	if len(parameters) > 0 {
		*self.Memory.Buffers[mem.Transmit] = []byte(parameters[0])
	}

	return Must(strconv.Atoi(parameters[0]))
}
