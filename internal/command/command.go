// ???: Move this and commands to script package?
package command

import (
	"fmt"
	"strconv"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

type Command struct {
	Memory *mem.Memory
}

func New(memory *mem.Memory) Command {
	return Command{Memory: memory}
}

func (c Command) Int(name string, parameters []string, index int, def int) int {
	if len(parameters) < index+1 || parameters[index] == "" {
		return def
	}

	defer etc.Recover(func(e error) { panic(fmt.Sprintf("invalid value: %s: %v", name, parameters[index])) })
	return etc.Must(strconv.Atoi(parameters[index]))
}
