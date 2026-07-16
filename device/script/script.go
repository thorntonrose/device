package script

import (
	"fmt"
	"log"
	"maps"
	"regexp"
	"slices"
	"strings"

	"github.com/thorntonrose/device/device/command"
	"github.com/thorntonrose/device/device/config"
	. "github.com/thorntonrose/device/device/etc"
	"github.com/thorntonrose/device/device/mem"
)

// syntax: ['*'|'+']<letter>[<parameter>(.<parameter>)*]
// A*B1+C1.-2.#3.'A1!'D.1E..2 => A, *B1, +C1.-2.#3.'A1!', D.1, E..2
var CommandPattern = regexp.MustCompile(`([*+]?[A-Z])((?:#\d|[-]?\d+|'[^']*'|)(?:\.(?:#\d|[-]?\d+|'[^']*'|))*)?`)

// multiple digit variables
// var CommandPattern = regexp.MustCompile(`([*+]?[A-Z])((?:[-#]?\d+|'[^']*'|)(?:\.(?:[-#]?\d+|'[^']*'|))*)?`)

type Script struct {
	Memory   *mem.Memory
	Commands map[string]Command
}

type Command interface {
	Run(parameters []string) (skip int)
}

func New(memory *mem.Memory) (script Script) {
	script.Memory = memory
	script.Commands = NewCommands(memory, &script)
	log.Printf("script.New: Commands: %v\n", script.CommandNames())

	return
}

func NewCommands(memory *mem.Memory, script *Script) map[string]Command {
	commands := make(map[string]Command)
	AddCommand(commands, "A", command.NewA(memory), 5)              // append data to dest
	AddCommand(commands, "+A", command.NewPlusA(memory), 1)         // compare variable
	AddCommand(commands, "B", command.NewB(memory), 1)              // set buffers
	AddCommand(commands, "D", command.NewD(memory), 5)              // delete from dest
	AddCommand(commands, "G", command.NewG(memory), 1)              // clear dest
	AddCommand(commands, "H", command.NewH(memory), 5)              // search in src
	AddCommand(commands, "I", command.NewI(memory), 1)              // compare
	AddCommand(commands, "+I", command.NewPlusI(memory), 1)         // send/receive
	AddCommand(commands, "O", command.NewO(memory), 5)              // move src pointer
	AddCommand(commands, "*O", command.NewStarO(memory), 1)         // do variable math
	AddCommand(commands, "P", command.NewP(memory), 5)              // display slot
	AddCommand(commands, "*Q", command.NewStarQ(memory), 1)         // set variable from buffer
	AddCommand(commands, "+Q", command.NewPlusQ(memory), 1)         // copy variable to dest
	AddCommand(commands, "R", command.NewR(memory), 5)              // append value to dest
	AddCommand(commands, "*L", command.NewStarL(memory, script), 5) // call
	AddCommand(commands, "*M", command.NewStarM(memory), 5)         // return
	AddCommand(commands, "*N", command.NewStarN(memory), 1)         // set variable
	AddCommand(commands, "T", command.NewT(memory), 5)              // do math on slot
	AddCommand(commands, "V", command.NewV(memory), 1)              // display buffer
	AddCommand(commands, "X", command.NewX(memory), 1)              // copy src to dest
	AddCommand(commands, "Y", command.NewY(memory), 1)              // append non-empty memory

	return commands
}

func AddCommand(commands map[string]Command, name string, command Command, minModel int) {
	if minModel <= config.Model {
		commands[name] = command
	}
}

//-----------------------------------------------------------------------------

func (self Script) Run(slotNum int) int {
	log.Printf("Run: %d\n", slotNum)
	return self.RunCommands(0, self.Parse(string(self.Memory.Slots[slotNum])))
}

func (self Script) RunCommands(index int, commands [][]string) int {
	for index >= 0 && index < len(commands) {
		index = self.Next(index, self.RunCommand(commands[index]))
	}

	return index
}

func (self Script) RunCommand(tokens []string) int {
	log.Printf("RunCommand: %v\n", tokens)
	runner := self.Commands[tokens[1]]
	Assert(runner != nil, fmt.Errorf("unknown command: %s", tokens[1]))

	return runner.Run(strings.Split(tokens[2], "."))
}

func (self Script) Parse(value string) [][]string {
	// expect: [[<match>, <modifier><letter>, <parameters>], ...]
	return CommandPattern.FindAllStringSubmatch(value, -1)
}

func (self Script) Next(index int, skip int) int {
	return If(skip == 0, index+1, If(skip > 0, index+1+skip, index+skip))
}

//-----------------------------------------------------------------------------

func (self Script) CommandNames() (names []string) {
	names = slices.Collect(maps.Keys(self.Commands))
	slices.SortFunc(names, func(a, b string) int { return self.CompareNames(a, b) })

	return
}

func (self Script) CompareNames(a, b string) int {
	a2 := strings.TrimLeft(a, "*+")
	b2 := strings.TrimLeft(b, "*+")

	return If(a2 < b2, -1, If(a2 > b2, 1, self.ComparePlus(a, b)))
}

func (self Script) ComparePlus(a, b string) int {
	aPlus := strings.HasPrefix(a, "+")
	bPlus := strings.HasPrefix(b, "+")

	return If(!aPlus && bPlus, -1, If(aPlus && !bPlus, 1, self.CompareStar(a, b)))
}

func (self Script) CompareStar(a, b string) int {
	aStar := strings.HasPrefix(a, "*")
	bStar := strings.HasPrefix(b, "*")

	return If(!aStar && bStar, -1, If(aStar && !bStar, 1, 0))
}

//-----------------------------------------------------------------------------

func (self Script) IsCommand(value string) bool {
	name := Value(self.Parse(value), [][]string{{"", ""}})[0][1]
	return self.Commands[name] != nil
}
