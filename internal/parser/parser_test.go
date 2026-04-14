package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/etc"
)

func TestBufSlot(t *testing.T) {
	assert.Equal(t, 1, BufSlot("001"))
}

func TestSlot_Invalid(t *testing.T) {
	assert.Equal(t, 1, BufSlot("1"))
	assert.PanicsWithValue(t, "invalid slot: A", func() { BufSlot("A") })
}

func TestText(t *testing.T) {
	assert.Equal(t, []byte(""), Text(""))
	assert.Equal(t, []byte(""), Text("; comment"))
	assert.Equal(t, []byte("HELLO"), Text("HELLO"))
	assert.Equal(t, []byte("HELLO"), Text("HELLO; comment"))
}

func TestAssignment(t *testing.T) {
	assertAssignment(t, 2, []byte(""), "002=")
	assertAssignment(t, 2, []byte("HELLO"), "002=HELLO")
}

func assertAssignment(t *testing.T, slot int, value []byte, line string) {
	s, v := Assignment(line, "=")
	assert.Equal(t, slot, s, line)
	assert.Equal(t, value, v, line)
}

//-----------------------------------------------------------------------------

func TestCommand(t *testing.T) {
	etc.Each([]string{"X", "+X", "*X", "X1", "X.1", "X1.1"}, func(command string) {
		assert.Equal(t, []byte(command), Command(command))
	})
}

func TestCommand_Invalid(t *testing.T) {
	assert.PanicsWithValue(t, "invalid command: _", func() { Command("_") })
}

func TestCommands(t *testing.T) {
	lines := []string{"X", "G"}

	index, slot, value := Commands(lines, 0, 2, []byte{})
	assert.Equal(t, len(lines), index)
	assert.Equal(t, 2, slot)
	assert.Equal(t, []byte("XG"), value)
}

func TestScript(t *testing.T) {
	assertScript(t, 0, 2, []byte(""), []string{"002$"})
	assertScript(t, 0, 2, []byte("X"), []string{"002$X"})
	assertScript(t, 0, 2, []byte("X"), []string{"002$", "X"})
}

func TestScript_Invalid(t *testing.T) {
	assert.PanicsWithValue(t, "invalid directive: A", func() { Script([]string{"A"}, 0) })
}

func assertScript(t *testing.T, lineNum int, slot int, value []byte, lines []string) {
	_, s, v := Script(lines, lineNum)
	assert.Equal(t, slot, s, lines[lineNum])
	assert.Equal(t, value, v, lines[lineNum])
}

//-----------------------------------------------------------------------------

func TestDirective(t *testing.T) {
	assertDirective(t, 0, 2, []byte("HELLO"), []string{"002=HELLO"})
	assertDirective(t, 0, 2, []byte("X"), []string{"002$X"})
}

func assertDirective(t *testing.T, lineNum int, slot int, value []byte, lines []string) {
	_, s, v := Directive(lines, lineNum)
	assert.Equal(t, slot, s, lines[lineNum])
	assert.Equal(t, value, v, lines[lineNum])
}

//-----------------------------------------------------------------------------

func TestStatement(t *testing.T) {
	assertStatement(t, 0, 0, nil, []string{""})
	assertStatement(t, 0, 0, nil, []string{"  "})
	assertStatement(t, 0, 0, nil, []string{"; comment"})
	assertStatement(t, 0, 2, []byte("HELLO"), []string{"002=HELLO"})
}

func assertStatement(t *testing.T, lineNum int, slot int, value []byte, lines []string) {
	_, s, v := Statement(lines, lineNum)
	assert.Equal(t, slot, s, lines[lineNum])
	assert.Equal(t, value, v, lines[lineNum])
}

//-----------------------------------------------------------------------------

func TestParse(t *testing.T) {
	program := "\n" +
		"; comment\n" +
		"002=HELLO\n" +
		"020$X\n"

	data := Parse(program)
	assert.Equal(t, map[int][]byte{2: []byte("HELLO"), 20: []byte("X")}, data)
}
