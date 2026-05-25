package script

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/thorntonrose/device/internal/command"
	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
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

	return
}

func NewCommands(memory *mem.Memory, script *Script) map[string]Command {
	return map[string]Command{
		"A":  command.NewA(memory),             // append data to dest
		"B":  command.NewB(memory),             // set buffers
		"D":  command.NewD(memory),             // delete from dest
		"G":  command.NewG(memory),             // clear dest
		"H":  command.NewH(memory),             // search in src
		"I":  command.NewI(memory),             // compare
		"O":  command.NewO(memory),             // move src pointer
		"P":  command.NewP(memory),             // display slot
		"+A": command.NewPlusA(memory),         // compare variable
		"+I": command.NewPlusI(memory),         // send/receive
		"+Q": command.NewPlusQ(memory),         // copy variable to dest
		"R":  command.NewR(memory),             // append value to dest
		"*L": command.NewStarL(memory, script), // call
		"*M": command.NewStarM(memory),         // return
		"*N": command.NewStarN(memory),         // set variable
		"*O": command.NewStarO(memory),         // do variable math
		"*Q": command.NewStarQ(memory),         // set variable from buffer
		"T":  command.NewT(memory),             // do math on slot
		"V":  command.NewV(memory),             // display buffer
		"X":  command.NewX(memory),             // copy src to dest
		"Y":  command.NewY(memory),             // append non-empty memory
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

func (self Script) IsCommand(value string) bool {
	name := Value(self.Parse(value), [][]string{{"", ""}})[0][1]
	return self.Commands[name] != nil
}
