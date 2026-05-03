package script

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// syntax: ['*'|'+']<letter>[<parameter>(.<parameter>)*]
// A*B1+C1.-2.#3.'A1!'D.1E..2 => A, *B1, +C1.-2.#3.'A1!', D.1, E..2
var CommandPattern = regexp.MustCompile(`([*+]?[A-Z])((?:#\d|[-]?\d+|'[^']*'|)(?:\.(?:#\d|[-]?\d+|'[^']*'|))*)?`)

// multiple digit variable numbers
// var CommandPattern = regexp.MustCompile(`([*+]?[A-Z])((?:[-#]?\d+|'[^']*'|)(?:\.(?:[-#]?\d+|'[^']*'|))*)?`)

type Script struct {
	Memory  *mem.Memory
	Runners map[string]Runner
}

type Runner interface {
	Run(parameters []string) (skip int)
}

func NewScript(memory *mem.Memory) Script {
	return Script{Memory: memory, Runners: NewRunners(memory)}
}

func NewRunners(memory *mem.Memory) map[string]Runner {
	return map[string]Runner{
		"+A": NewPlusA(memory), // compare variable
		"B":  NewB(memory),     // set buffers
		"G":  NewG(memory),     // clear dest
		"I":  NewI(memory),     // compare
		"+I": NewPlusI(memory), // send/receive
		"*N": NewStarN(memory), // set variable
		"O":  NewO(memory),     // move src pointer
		"*O": NewStarO(memory), // variable math
		"P":  NewP(memory),     // display slot
		"V":  NewV(memory),     // display buffer
		"X":  NewX(memory),     // copy src to dest
		"Y":  NewY(memory),     // dump memory
	}
}

//-----------------------------------------------------------------------------

func (self Script) Run(slotNum int) int {
	log.Printf("Script.Run: slot: %d\n", slotNum)
	return self.RunCommands(0, self.Commands(string(self.Memory.Slots[slotNum])))
}

func (self Script) RunCommands(index int, commands [][]string) int {
	log.Printf("Script.RunCommands: %v\n", commands)

	for index >= 0 && index < len(commands) {
		index = self.Next(index, self.RunCommand(commands[index]))
	}

	return index
}

func (self Script) RunCommand(tokens []string) int {
	log.Printf("Script.RunCommand: tokens: %v\n", tokens)
	runner := self.Runners[tokens[1]]
	etc.Assert(runner != nil, fmt.Sprintf("unknown command: %s", tokens[1]))

	return runner.Run(strings.Split(tokens[2], "."))
}

func (self Script) Commands(value string) [][]string {
	// expect: [[<match>, <modifier><letter>, <parameters>], ...]
	return CommandPattern.FindAllStringSubmatch(value, -1)
}

func (self Script) Next(index int, skip int) int {
	return etc.If(skip == 0, index+1, etc.If(skip > 0, index+1+skip, index+skip))
}

//-----------------------------------------------------------------------------

func (self Script) IsCommand(value string) bool {
	name := etc.Value(self.Commands(value), [][]string{{"", ""}})[0][1]
	return self.Runners[name] != nil
}
