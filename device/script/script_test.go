//go:build test

package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/device/mem"
	_ "github.com/thorntonrose/device/device/testing"
)

func TestNext(t *testing.T) {
	script := New(mem.New())
	assert.Equal(t, 1, script.Next(0, 0))
	assert.Equal(t, 2, script.Next(0, 1))
	assert.Equal(t, -1, script.Next(0, -1))
}

func TestParse(t *testing.T) {
	memory := mem.New()
	memory.Set(20, []byte("A*B1+C1.-2.#3.'A1!'D.1E..2"))

	commands := New(memory).Parse(string(memory.Slots[20]))
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
	memory.Set(mem.Receive+1, []byte("FOO"))

	script := NewTestScript(memory)
	assert.Equal(t, 1, script.RunCommand([]string{"Z1.1", "Z", "1.1"}))
	assert.Equal(t, []byte("1"), *memory.Buffers[mem.Transmit])
}

//-----------------------------------------------------------------------------

func TestRun(t *testing.T) {
	AssertRun(t, "Z0Z0", 2, "0")
	AssertRun(t, "Z1Z", 2, "1")
	AssertRun(t, "Z-1Z", -1, "-1")
}

func AssertRun(t *testing.T, text string, expectedIndex int, expectedData string) {
	script := NewTestScript(mem.New())
	script.Memory.Set(20, []byte(text))

	index := script.Run(20)
	assert.Equal(t, expectedIndex, index)
	assert.Equal(t, []byte(expectedData), *script.Memory.Buffers[mem.Transmit])
}

//-----------------------------------------------------------------------------

func TestRun_Subroutine(t *testing.T) {
	script := New(mem.New())
	script.Commands["Z"] = NewZ(script.Memory)
	script.Memory.Set(20, []byte("*L21"))
	script.Memory.Set(21, []byte("Z0*MZ"))

	script.Run(20)
	assert.Equal(t, []byte("0"), *script.Memory.Buffers[mem.Transmit])
}

//-----------------------------------------------------------------------------

func TestIsCommand(t *testing.T) {
	script := NewTestScript(mem.New())

	for _, command := range []string{"Z", "+Z", "*Z", "Z1", "Z.1", "Z1.1"} {
		assert.True(t, script.IsCommand(command), command)
	}

	for _, nonCommand := range []string{"_", "A!", "A#10"} {
		assert.False(t, script.IsCommand(nonCommand), nonCommand)
	}
}
