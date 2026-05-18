package parser

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/script"
)

const CommentMarker = ";"

var SingleCommandPattern = regexp.MustCompile(`^` + script.CommandPattern.String() + `$`)

type Parser struct {
	Script script.Script
}

func New(script script.Script) Parser {
	return Parser{Script: script}
}

//-----------------------------------------------------------------------------

func (self Parser) Parse(program string) (data map[int][]byte) {
	log.Println("parser.Parse")
	data = map[int][]byte{}
	self.ParseLines(data, strings.Split(program, "\n"), 0)

	delete(data, 0)
	return
}

func (self Parser) ParseLines(data map[int][]byte, lines []string, index int) int {
	log.Printf("parser.ParseLines: line: %d, %s\n", index, lines[index])

	for index < len(lines) {
		index = self.ParseLine(data, lines, index)
	}

	return index
}

func (self Parser) ParseLine(data map[int][]byte, lines []string, index int) (newIndex int) {
	newIndex, slotNum, value := self.Statement(lines, index)

	if value != nil {
		log.Printf("parser.ParseLine: newIndex: %d, slotNum: %d, value: %s\n", newIndex, slotNum, value)
		Assert(len(data[slotNum]) == 0, fmt.Errorf("duplicate slot: '%d'", slotNum))
		data[slotNum] = value
	}

	return
}

//-----------------------------------------------------------------------------

func (self Parser) Statement(lines []string, index int) (int, int, []byte) {
	if self.IsBlank(lines[index]) || self.IsComment(lines[index]) {
		return index + 1, 0, nil
	}

	return self.Directive(lines, index)
}

func (self Parser) IsBlank(line string) bool {
	return strings.TrimSpace(line) == ""
}

func (self Parser) IsComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), CommentMarker)
}

func (self Parser) Directive(lines []string, index int) (int, int, []byte) {
	if slot, value := self.DataDirective(lines[index], "="); value != nil {
		return index + 1, slot, value
	}

	return self.ScriptDirective(lines, index)
}

//-----------------------------------------------------------------------------

func (self Parser) DataDirective(line string, sep string) (int, []byte) {
	if tokens := strings.SplitN(line, sep, 2); len(tokens) == 2 {
		return self.SlotNum(tokens[0]), self.Text(tokens[1])
	}

	return 0, nil
}

func (self Parser) SlotNum(token string) int {
	defer Recover(func(e error) { panic(fmt.Sprintf("invalid slot: '%s'", token)) })
	return Must(strconv.Atoi(token))
}

func (self Parser) Text(token string) []byte {
	index := strings.Index(token, CommentMarker)
	return []byte(token[:If(index == -1, len(token), index)])
}

//-----------------------------------------------------------------------------

func (self Parser) ScriptDirective(lines []string, index int) (int, int, []byte) {
	line := lines[index]

	if slot, value := self.DataDirective(line, "$"); value != nil {
		return self.Commands(lines, index+1, slot, self.Command(string(value)))
	}

	panic(fmt.Sprintf("invalid directive: '%s'", line))
}

func (self Parser) Commands(lines []string, index int, slotNum int, value []byte) (newIndex int, currSlotNum int,
	newValue []byte,
) {
	defer Recover(func(e error) {
		newIndex, currSlotNum, newValue = self.CommandsRecover(e, index, slotNum, value)
	})

	for index < len(lines) {
		value = append(value, self.Command(string(self.Text(lines[index])))...)
		index++
	}

	return index, slotNum, value
}

func (self Parser) CommandsRecover(e error, index int, slotNum int, value []byte) (int, int, []byte) {
	Assert(strings.Contains(e.Error(), "invalid command"), e)
	return index, slotNum, value
}

func (self Parser) Command(line string) []byte {
	line = strings.TrimSpace(line)

	if line != "" {
		Assert(SingleCommandPattern.MatchString(line), fmt.Errorf("invalid command: '%s'", line))
		Assert(self.Script.IsCommand(line), fmt.Errorf("unknown command: '%s'", line))
	}

	return []byte(line)
}
