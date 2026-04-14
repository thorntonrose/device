package script

import (
	"regexp"
	"strings"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

// syntax: ['*'|'+']<letter>[<parameter>(.<parameter>)*]
// A*B1+C1.-2.#3.'A1!'D.1E..2 => A, *B1, +C1.-2.#3.'A1!', D.1, E..2
var CommandPattern = regexp.MustCompile(`([*+]?[A-Z])((?:[-#]?\d+|'[^']*'|)(?:\.(?:[-#]?\d+|'[^']*'|))*)?`)

type Script struct {
	Memory  mem.Memory
	Runners map[string]Runner
}

type Runner interface {
	Run(parameters []string) (skip int)
}

func New(memory mem.Memory, runners map[string]Runner) Script {
	return Script{Memory: memory, Runners: runners}
}

func (s Script) Run(slot int) int {
	commands := s.Commands(slot)
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
	return s.Runners[tokens[1]].Run(strings.Split(tokens[2], "."))
}
