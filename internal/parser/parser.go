package parser

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/script"
)

const (
	CommentMarker = ";"
	SlotLength    = 3
)

var SingleCommandPattern = regexp.MustCompile(`^` + script.CommandPattern.String() + `$`)

type Parser struct {
	Script script.Script
}

func New(script script.Script) Parser {
	return Parser{Script: script}
}

// Syntax:
//
// <program> ::= <statement>+
// <statement> ::= <comment> | <data-directive> | <script-directive>
// <comment> ::= ';'[<text>]<newline>
//
// <data-directive> ::= <slot>'='<text><eol>
// <slot> ::= ['0']*<digit>+
//
// <script-directive> ::= <slot>'$'(<space>*<command>[<eol>])+
// <command> ::= ['+' | '*']<'A'..'Z'>[<parameters>]
// <parameters> ::= <parameter>('.'<parameter>)*
// <parameter> ::= '#'<digit> | ['-']<integer> | "'"<text>"'"
//
// <eol> ::= <comment> | <newline>
//
// Example:
//
// ; program
// 002=FOO
// 003=123
// 020$X              ; src -> dest
// 021$Y1.-2.#3.'A1!' ; not a real command
func (self Parser) Parse(program string) map[int][]byte {
	log.Println("parser.Parse")
	data := map[int][]byte{}
	lines := strings.Split(program, "\n")
	index := 0

	for index < len(lines) {
		log.Printf("parser.Parse: line: %d, %s\n", index, lines[index])

		newIndex, slotNum, value := self.Statement(lines, index)
		log.Printf("parser.Parse: newIndex: %d, slot: %d, value: %s\n", newIndex, slotNum, value)

		data[slotNum] = value
		index = newIndex
	}

	delete(data, 0)
	return data
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
	defer etc.Recover(func(e error) { panic("invalid slot: " + token) })
	return etc.Must(strconv.Atoi(etc.Value(token, "0")))
}

func (self Parser) Text(token string) []byte {
	index := strings.Index(token, CommentMarker)
	return []byte(token[:etc.If(index == -1, len(token), index)])
}

//-----------------------------------------------------------------------------

func (self Parser) ScriptDirective(lines []string, index int) (int, int, []byte) {
	line := lines[index]

	if slot, value := self.DataDirective(line, "$"); value != nil {
		return self.Commands(lines, index+1, slot, self.Command(string(value)))
	}

	panic("invalid directive: " + line)
}

func (self Parser) Commands(lines []string, index int, slotNum int, value []byte) (newIndex int, currSlotNum int, newValue []byte) {
	defer etc.Recover(func(e error) { newIndex, currSlotNum, newValue = self.CommandsRecover(e, index, slotNum, value) })

	for index < len(lines) {
		value = append(value, self.Command(string(self.Text(lines[index])))...)
		index++
	}

	return index, slotNum, value
}

func (self Parser) CommandsRecover(e error, index int, slotNum int, value []byte) (int, int, []byte) {
	etc.Assert(strings.Contains(e.Error(), "invalid command"), e.Error())
	return index, slotNum, value
}

func (self Parser) Command(line string) []byte {
	line = strings.TrimSpace(line)

	if line != "" {
		etc.Assert(SingleCommandPattern.MatchString(line), fmt.Sprintf("invalid command: %s", line))
		etc.Assert(self.Script.IsCommand(line), fmt.Sprintf("unknown command: %s", line))
	}

	return []byte(line)
}
