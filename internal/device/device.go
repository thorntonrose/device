package device

import (
	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/mem"
	"github.com/thorntonrose/device/internal/parser"
	"github.com/thorntonrose/device/internal/script"
)

type Device struct {
	Memory  *mem.Memory
	Runners map[string]script.Runner
	Script  script.Script
}

func New() Device {
	d := Device{Memory: mem.New()}
	d.Runners = NewCommands(d.Memory)
	d.Script = script.New(d.Memory, d.Runners)

	return d
}

func NewCommands(memory *mem.Memory) map[string]script.Runner {
	return map[string]script.Runner{
		"+I": command.NewPlusI(memory),
		"X":  command.NewX(memory),
		"Y":  command.NewY(memory),
	}
}

//-----------------------------------------------------------------------------

func (self Device) Load(program string) {
	self.Memory.Load(parser.Parse(program))
}

func (self Device) Run(slotNum int) {
	self.Script.Run(slotNum)
}
