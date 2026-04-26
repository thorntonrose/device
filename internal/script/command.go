// ???: Move this and commands to script package?
package script

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

type Command struct {
	Memory *mem.Memory
}

func NewCommand(memory *mem.Memory) Command {
	return Command{Memory: memory}
}

//-----------------------------------------------------------------------------

func (self Command) Bytes(name string, parameters []string, index int, def []byte) []byte {
	if len(parameters) < index+1 || parameters[index] == "" {
		return def
	}

	return self.BytesFromParameter(name, parameters[index])
}

func (self Command) BytesFromParameter(name string, parameter string) []byte {
	if parameter[0] == '\'' {
		return []byte(parameter[1 : len(parameter)-1])
	}

	return []byte{byte(self.Range(name, []string{parameter}, 0, 0, 0, 255))}

}

func (self Command) Code(name string, parameters []string, index int, def int, values []int) int {
	value := self.Int(name, parameters, index, def)
	etc.Assert(slices.Contains(values, value), fmt.Sprintf("%s: %d not in %v", name, value, values))

	return value
}

func (self Command) NonNegative(name string, parameters []string, index int, def int) int {
	value := self.Int(name, parameters, index, def)
	etc.Assert(value >= 0, fmt.Sprintf("%s: %d < 0", name, value))

	return value
}

func (self Command) Positive(name string, parameters []string, index int, def int) int {
	value := self.Int(name, parameters, index, def)
	etc.Assert(value > 0, fmt.Sprintf("%s: %d <= 0", name, value))

	return value
}

func (self Command) Range(name string, parameters []string, index int, def int, min int, max int) int {
	value := self.Int(name, parameters, index, def)
	etc.Assert(value >= min && value <= max, fmt.Sprintf("%s: %d not in %d - %d", name, value, min, max))

	return value
}

func (self Command) Variable(name string, parameters []string, index int, def int) int {
	if len(parameters) < index+1 || parameters[index] == "" {
		return def
	}

	parameter := parameters[index]
	etc.Assert(parameter[0] == '#', fmt.Sprintf("%s: %s not a variable", name, parameter))

	return self.VarNum(name, parameter)
}

//-----------------------------------------------------------------------------

func (self Command) Int(name string, parameters []string, index int, def int) int {
	if len(parameters) < index+1 || parameters[index] == "" {
		return def
	}

	return self.IntFromParameter(name, parameters[index])
}

func (self Command) IntFromParameter(name string, parameter string) int {
	if parameter[0] == '#' {
		return self.Memory.Variables[self.VarNum(name, parameter)]
	}

	return self.IntFromValue(name, parameter)
}

func (self Command) VarNum(name string, parameter string) int {
	varNum := self.IntFromValue(name, parameter[1:])
	etc.Assert(varNum >= 0 && varNum < mem.MaxVariables, fmt.Sprintf("%s: invalid variable: %d", name, varNum))

	return varNum
}

func (self Command) IntFromValue(name string, parameter string) int {
	defer etc.Recover(func(e error) { panic(fmt.Sprintf("%s: %s not integer", name, parameter)) })
	return etc.Must(strconv.Atoi(parameter))
}

// Model 5?
//
// func (self Command) IntFromSlot(name string, parameter string) int {
// 	slotNum := self.IntFromValue(name, parameter[1:])
// 	data := self.Memory.Slots[slotNum]
//
// 	defer etc.Recover(func(e error) {
// 		panic(fmt.Sprintf("%s: slot number: %d: invalid value: %s", name, slotNum, string(data)))
// 	})
//
// 	return etc.Must(strconv.Atoi(string(data)))
// }
