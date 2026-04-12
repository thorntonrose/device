package script

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

var CommandPattern = regexp.MustCompile(`(\*|\+)?([A-Z])([0-9\.])*`)

type Script struct {
	Memory   *mem.Memory
	Commands map[string]command.Runner
}

func New(memory *mem.Memory, commands map[string]command.Runner) Script {
	return Script{Memory: memory, Commands: commands}
}

func (s Script) Run(slot int) {
	etc.Each(CommandPattern.FindAllStringSubmatch(string(s.Memory.Get(slot)), -1), func(tokens []string) {
		s.RunCommand(tokens[1]+tokens[2], tokens[3])
	})
}

func (s Script) RunCommand(name string, token string) {
	s.Commands[name].Run(s.ToParameters(token))
}

func (s Script) ToParameters(token string) (parameters command.Parameters) {
	etc.EachWithIndex(strings.Split(token, "."), func(val string, i int) { parameters[i] = s.ToInt(val) })
	return parameters
}

func (s Script) ToInt(val string) int {
	defer etc.Recover(func(e error) { panic("invalid parameter: " + val) })
	return etc.Must(strconv.Atoi(etc.Value(val, "0")))
}
