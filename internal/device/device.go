package device

import (
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
	d.Script = script.NewScript(d.Memory)

	return d
}

//-----------------------------------------------------------------------------

func (self Device) Load(program string) {
	self.Memory.Load(parser.Parse(program))
}

func (self Device) Run(slotNum int) {
	self.Script.Run(slotNum)
}
