package script

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

func TestNext(t *testing.T) {
	script := New(mem.New(), nil)
	assert.Equal(t, 1, script.Next(0, 0))
	assert.Equal(t, 2, script.Next(0, 1))
	assert.Equal(t, -1, script.Next(0, -1))
}

func TestCommands(t *testing.T) {
	memory := mem.New()
	memory.Set(20, []byte("A*B1+C1.-2.#3.'A1!'D.1E..2"))

	commands := New(memory, nil).Commands(20)
	assert.Equal(t, [][]string{
		{"A", "A", ""},
		{"*B1", "*B", "1"},
		{"+C1.-2.#3.'A1!'", "+C", "1.-2.#3.'A1!'"},
		{"D.1", "D", ".1"},
		{"E..2", "E", "..2"},
	}, commands)
}

func TestRunCommand(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Slot(mem.Receive), []byte("HELLO"))

	script := New(memory, map[string]Runner{"S": NewS(memory)})
	assert.Equal(t, 1, script.RunCommand([]string{"S1.1", "S", "1.1"}))
	assert.Equal(t, []byte("1"), memory.Get(mem.Slot(mem.Transmit)))
}

func TestRun(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Slot(mem.Receive), []byte("HELLO"))
	memory.Set(20, []byte("S0S0"))

	index := New(memory, map[string]Runner{"S": NewS(memory)}).Run(20)
	assert.Equal(t, 2, index)
}

func TestRun_Skip(t *testing.T) {
	AssertSkip(t, "S1X", 2)
	AssertSkip(t, "S-1X", -1)
}

func AssertSkip(t *testing.T, text string, expectedIndex int) {
	memory := mem.New()
	memory.Set(mem.Slot(mem.Receive), []byte("HELLO"))
	memory.Set(20, []byte(text))

	index := New(memory, map[string]Runner{"S": NewS(memory)}).Run(20)
	assert.Equal(t, []byte{}, memory.Get(mem.Slot(mem.Transmit)))
	assert.Equal(t, expectedIndex, index)
}

//-----------------------------------------------------------------------------

type S struct {
	command.Command
}

func NewS(memory *mem.Memory) S {
	return S{command.Command{Memory: memory}}
}

func (s S) Run(parameters []string) int {
	if len(parameters) > 1 {
		s.Memory.Set(mem.Slot(mem.Transmit), []byte(parameters[1]))
	}

	return etc.Must(strconv.Atoi(parameters[0]))
}
