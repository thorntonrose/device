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

func New(memory *mem.Memory, runners map[string]Runner) Script {
	return Script{Memory: memory, Runners: runners}
}

func (self Script) Run(slotNum int) int {
	log.Printf("Script.Run: slot: %d", slotNum)
	commands := self.Commands(slotNum)
	log.Printf("Script.Run: commands: %v", commands)
	index := 0

	for index >= 0 && index < len(commands) {
		index = self.Next(index, self.RunCommand(commands[index]))
	}

	return index
}

func (self Script) Commands(slotNum int) [][]string {
	// expect: [<match>, <modifier><letter>, <parameters>]
	return CommandPattern.FindAllStringSubmatch(string(self.Memory.Slots[slotNum]), -1)
}

func (self Script) Next(index int, skip int) int {
	return etc.If(skip == 0, index+1, etc.If(skip > 0, index+1+skip, index+skip))
}

func (self Script) RunCommand(tokens []string) int {
	log.Printf("Script.RunCommand: tokens: %v", tokens)
	return self.Runners[tokens[1]].Run(strings.Split(tokens[2], "."))
}
