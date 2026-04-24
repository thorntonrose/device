package device

import (
	"log"

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

func (d Device) Load(program string) {
	log.Println("Device.Load ...")
	d.Memory.Load(parser.Parse(program))
}

func (d Device) Run(slot int) {
	log.Printf("Device.Run: slot: %d", slot)
	d.Script.Run(slot)
}
