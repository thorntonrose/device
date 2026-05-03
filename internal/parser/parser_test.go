package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/config"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
	"github.com/thorntonrose/device/internal/script"
)

func TestMain(m *testing.M) {
	defer config.InitLogger()()
	m.Run()
}

//-----------------------------------------------------------------------------

func TestSlotNum(t *testing.T) {
	p := NewTestParser()
	assert.Equal(t, 1, p.SlotNum("001"))
	assert.Equal(t, 1, p.SlotNum("1"))
	assert.PanicsWithValue(t, "invalid slot: A", func() { p.SlotNum("A") })
}

func TestText(t *testing.T) {
	p := NewTestParser()
	assert.Equal(t, []byte(""), p.Text(""))
	assert.Equal(t, []byte(""), p.Text("; comment; ..."))
	assert.Equal(t, []byte("FOO"), p.Text("FOO"))
	assert.Equal(t, []byte("FOO"), p.Text("FOO; comment"))
}

func TestDataDirective(t *testing.T) {
	assertDataDirective(t, 2, []byte(""), "002=")
	assertDataDirective(t, 2, []byte("FOO"), "002=FOO")
}

func assertDataDirective(t *testing.T, slotNum int, value []byte, line string) {
	s, v := NewTestParser().DataDirective(line, "=")
	assert.Equal(t, slotNum, s, line)
	assert.Equal(t, value, v, line)
}

//-----------------------------------------------------------------------------

func TestCommand(t *testing.T) {
	p := NewTestParser()

	etc.Each([]string{"Z", "+Z", "*Z", "Z1", "Z.1", "Z1.1"}, func(command string) {
		assert.Equal(t, []byte(command), p.Command(command))
	})
}

func TestCommand_Invalid(t *testing.T) {
	p := NewTestParser()
	assert.PanicsWithValue(t, "invalid command: _", func() { p.Command("_") })
	assert.PanicsWithValue(t, "invalid command: A!", func() { p.Command("A!") })
	assert.PanicsWithValue(t, "invalid command: A#12", func() { p.Command("A#12") })
}

func TestCommands(t *testing.T) {
	lines := []string{"Z", "+Z"}

	index, slot, value := NewTestParser().Commands(lines, 0, 2, []byte{})
	assert.Equal(t, len(lines), index)
	assert.Equal(t, 2, slot)
	assert.Equal(t, []byte("Z+Z"), value)
}

func TestScriptDirective(t *testing.T) {
	assertScriptDirective(t, 0, 2, []byte(""), []string{"002$"})
	assertScriptDirective(t, 0, 2, []byte("Z"), []string{"002$Z"})
	assertScriptDirective(t, 0, 2, []byte("Z"), []string{"002$", "Z"})
}

func assertScriptDirective(t *testing.T, lineNum int, slotNum int, value []byte, lines []string) {
	_, s, v := NewTestParser().ScriptDirective(lines, lineNum)
	assert.Equal(t, slotNum, s, lines[lineNum])
	assert.Equal(t, value, v, lines[lineNum])
}

func TestScriptDirective_Invalid(t *testing.T) {
	assert.PanicsWithValue(t, "invalid directive: A", func() { NewTestParser().ScriptDirective([]string{"A"}, 0) })
}

//-----------------------------------------------------------------------------

func TestDirective(t *testing.T) {
	assertDirective(t, 0, 2, []byte("FOO"), []string{"002=FOO"})
	assertDirective(t, 0, 2, []byte("Z"), []string{"002$Z"})
}

func assertDirective(t *testing.T, lineNum int, slotNum int, value []byte, lines []string) {
	_, s, v := NewTestParser().Directive(lines, lineNum)
	assert.Equal(t, slotNum, s, lines[lineNum])
	assert.Equal(t, value, v, lines[lineNum])
}

//-----------------------------------------------------------------------------

func TestStatement(t *testing.T) {
	assertStatement(t, 0, 0, nil, []string{""})
	assertStatement(t, 0, 0, nil, []string{"  "})
	assertStatement(t, 0, 0, nil, []string{"; comment"})
	assertStatement(t, 0, 2, []byte("FOO"), []string{"002=FOO"})
}

func assertStatement(t *testing.T, lineNum int, slotNum int, value []byte, lines []string) {
	_, s, v := NewTestParser().Statement(lines, lineNum)
	assert.Equal(t, slotNum, s, lines[lineNum])
	assert.Equal(t, value, v, lines[lineNum])
}

func TestParse(t *testing.T) {
	program := strings.Join([]string{
		"; test",
		"002=FOO",
		"020$Z    ; do z",
		"021$Z1.2 ; do z with 1.2",
	}, "\n")

	data := NewTestParser().Parse(program)
	assert.Equal(t, map[int][]byte{2: []byte("FOO"), 20: []byte("Z"), 21: []byte("Z1.2")}, data)
}

//-----------------------------------------------------------------------------

func NewTestParser() Parser {
	return New(script.NewTestScript(mem.New()))
}
