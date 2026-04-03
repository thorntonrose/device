package script

import (
	"github.com/thorntonrose/device/internal/buf"
	"github.com/thorntonrose/device/internal/mem"
)

const MaxVariables = 10

type Script struct {
	Memory    mem.Memory
	BufferSet buf.BufferSet
	Variables [MaxVariables][]byte
}

func NewScript(memory mem.Memory, bufferSet buf.BufferSet) Script {
	return Script{Memory: memory, BufferSet: bufferSet, Variables: [MaxVariables][]byte{}}
}

func (s Script) Run(location int) {
	// ...
}

//-----------------------------------------------------------------------------

type Command struct {
	Script     Script
	Parameters []string
}
