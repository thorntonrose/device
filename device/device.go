package device

import (
	"github.com/thorntonrose/device/device/mem"
	"github.com/thorntonrose/device/device/parser"
	"github.com/thorntonrose/device/device/script"
)

type Device struct {
	Memory *mem.Memory
	Script script.Script
}

func New() (d Device) {
	d.Memory = mem.New()
	d.Script = script.New(d.Memory)

	return
}

//-----------------------------------------------------------------------------

func (self Device) Load(program string) {
	self.Memory.Load(parser.New(self.Script).Parse(program))
}

func (self Device) Run() {
	self.Script.Run(20)
}
