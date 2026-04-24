package script

import (
	"log"
	"regexp"
	"strings"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// syntax: ['*'|'+']<letter>[<parameter>(.<parameter>)*]
// A*B1+C1.-2.#3.'A1!'D.1E..2 => A, *B1, +C1.-2.#3.'A1!', D.1, E..2
var CommandPattern = regexp.MustCompile(`([*+]?[A-Z])((?:[-#]?\d+|'[^']*'|)(?:\.(?:[-#]?\d+|'[^']*'|))*)?`)

type Script struct {
	Memory  *mem.Memory
	Runners map[string]Runner
}

type Runner interface {
	Run(parameters []string) (skip int)
}

func NewScript(memory *mem.Memory) Script {
	return Script{Memory: memory, Runners: NewCommands(memory)}
}

func NewCommands(memory *mem.Memory) map[string]Runner {
	return map[string]Runner{
		"X": NewX(memory),
	}
}

//-----------------------------------------------------------------------------

func (s Script) Run(slot int) int {
	log.Printf("Script.Run: slot: %d", slot)
	commands := s.Commands(slot)
	log.Printf("Script.Run: commands: %v", commands)
	index := 0

	for index >= 0 && index < len(commands) {
		index = s.Next(index, s.RunCommand(commands[index]))
	}

	return index
}

func (s Script) Commands(slot int) [][]string {
	// expect: [<match>, <modifier><letter>, <parameters>]
	return CommandPattern.FindAllStringSubmatch(string(s.Memory.Get(slot)), -1)
}

func (s Script) Next(index int, skip int) int {
	return etc.If(skip == 0, index+1, etc.If(skip > 0, index+1+skip, index+skip))
}

func (s Script) RunCommand(tokens []string) int {
	log.Printf("Script.RunCommand: tokens: %v", tokens)
	return s.Runners[tokens[1]].Run(strings.Split(tokens[2], "."))
}
