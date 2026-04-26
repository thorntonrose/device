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

// Syntax:
//
// <program> ::= <statement>+
// <statement> ::= (<comment> | <assignment> | <script>)+
// <comment> ::= ';'[<text>]<newline>
//
// <assignment> ::= <slot>'='<text><eol>
// <slot> ::= ['0']*<digit>+
//
// <script> ::= (<slot>'$'(<space>*<command>[<eol>])+
// <command> ::= ['+' | '*']<'A'..'Z'>[<parameters>]
// <parameters> ::= <parameter>('.'<parameter>)*
// <parameter> ::= '#'<digit> | ['-']<integer> | "'"<text>"'"
//
// <eol> ::= <comment> | <newline>
//
// Example:
//
// 002=FOO
// 003=123
// 020$X
// 021$X1.-2.#3.'A1!' ; not a real command
func Parse(program string) map[int][]byte {
	log.Printf("Parse ...")
	data := map[int][]byte{}
	lines := strings.Split(program, "\n")
	index := 0

	for index < len(lines) {
		log.Printf("parser.Parse: line: %d, %s", index, lines[index])

		newIndex, slotNum, value := Statement(lines, index)
		log.Printf("parser.Parse: slot: %d, value: %s", slotNum, value)

		data[slotNum] = value
		index = newIndex
	}

	delete(data, 0)
	return data
}

//-----------------------------------------------------------------------------

func Statement(lines []string, index int) (int, int, []byte) {
	if IsBlank(lines[index]) || IsComment(lines[index]) {
		return index + 1, 0, nil
	}

	return Directive(lines, index)
}

func IsBlank(line string) bool {
	return strings.TrimSpace(line) == ""
}

func IsComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), CommentMarker)
}

func Directive(lines []string, index int) (int, int, []byte) {
	if slot, value := Assignment(lines[index], "="); value != nil {
		return index + 1, slot, value
	}

	return Script(lines, index)
}

//-----------------------------------------------------------------------------

func Assignment(line string, sep string) (int, []byte) {
	if tokens := strings.SplitN(line, sep, 2); len(tokens) == 2 {
		return BufSlotNum(tokens[0]), Text(tokens[1])
	}

	return 0, nil
}

func BufSlotNum(token string) int {
	defer etc.Recover(func(e error) { panic("invalid slot: " + token) })
	return etc.Must(strconv.Atoi(etc.Value(token, "0")))
}

func Text(token string) []byte {
	index := strings.LastIndex(token, CommentMarker)
	return []byte(token[:etc.If(index == -1, len(token), index)])
}

//-----------------------------------------------------------------------------

func Script(lines []string, index int) (int, int, []byte) {
	line := lines[index]

	if slot, value := Assignment(line, "$"); value != nil {
		return Commands(lines, index+1, slot, Command(string(value)))
	}

	panic("invalid directive: " + line)
}

func Commands(lines []string, index int, slotNum int, value []byte) (int, int, []byte) {
	for index < len(lines) {
		value = append(value, Command(string(Text(lines[index])))...)
		index++
	}

	return index, slotNum, value
}

func Command(line string) []byte {
	line = strings.TrimSpace(line)
	etc.Assert(line == "" || SingleCommandPattern.MatchString(line), fmt.Sprintf("invalid command: %s", line))

	return []byte(line)
}
