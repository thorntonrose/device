package script

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/config"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

func TestMain(m *testing.M) {
	defer config.InitLogger()()
	m.Run()
}

//-----------------------------------------------------------------------------

func TestNext(t *testing.T) {
	script := NewScript(mem.New())
	assert.Equal(t, 1, script.Next(0, 0))
	assert.Equal(t, 2, script.Next(0, 1))
	assert.Equal(t, -1, script.Next(0, -1))
}

func TestCommands(t *testing.T) {
	memory := mem.New()
	memory.Set(20, []byte("A*B1+C1.-2.#3.'A1!'D.1E..2"))

	commands := NewScript(memory).Commands(20)
	assert.Equal(t, [][]string{{"A", "A", ""}, {"*B1", "*B", "1"}, {"+C1.-2.#3.'A1!'", "+C", "1.-2.#3.'A1!'"},
		{"D.1", "D", ".1"}, {"E..2", "E", "..2"}}, commands)
}

func TestRunCommand(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Receive+1, []byte("FOO"))

	script := NewTestScript(memory)
	assert.Equal(t, 1, script.RunCommand([]string{"Z1.1", "Z", "1.1"}))
	assert.Equal(t, []byte("1"), *memory.Buffers[mem.Transmit])
}

func TestRun(t *testing.T) {
	memory := mem.New()
	*memory.Buffers[mem.Receive] = []byte("FOO")
	memory.Set(20, []byte("Z0Z0"))

	index := NewTestScript(memory).Run(20)
	assert.Equal(t, 2, index)
}

func TestRun_Skip(t *testing.T) {
	AssertSkip(t, "Z1X", 2)
	AssertSkip(t, "Z-1X", -1)
}

func AssertSkip(t *testing.T, text string, expectedIndex int) {
	memory := mem.New()
	*memory.Buffers[mem.Receive] = []byte("FOO")
	memory.Set(20, []byte(text))

	index := NewTestScript(memory).Run(20)
	assert.Equal(t, []byte{}, *memory.Buffers[mem.Transmit])
	assert.Equal(t, expectedIndex, index)
}

//-----------------------------------------------------------------------------

func NewTestScript(memory *mem.Memory) Script {
	script := NewScript(memory)
	script.Runners = map[string]Runner{"Z": NewZ(memory)}

	return script
}

type Z struct {
	Command
}

func NewZ(memory *mem.Memory) Z {
	return Z{NewCommand(memory)}
}

func (self Z) Run(parameters []string) int {
	if len(parameters) > 1 {
		*self.Memory.Buffers[mem.Transmit] = []byte(parameters[1])
	}

	return etc.Must(strconv.Atoi(parameters[0]))
}
