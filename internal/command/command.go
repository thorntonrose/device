package command

import (
	"fmt"
	"strconv"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

type Command struct {
	Memory mem.Memory
}

func New(memory mem.Memory) Command {
	return Command{Memory: memory}
}

func (c Command) ToInt(name string, parameters []string, index int, def int) int {
	defer etc.Recover(func(e error) { panic(fmt.Sprintf("invalid %s: %v", name, parameters)) })
	defString := fmt.Sprintf("%d", def)

	return etc.If(len(parameters) > index, etc.Must(strconv.Atoi(etc.Value(parameters[index], defString))), def)
}
