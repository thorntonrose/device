package device

import (
	"fmt"
	"strings"
)

const (
	MaxMemory              = 40
	TransmitBufferLocation = 2
	ReceiveBufferLocation  = 3
)

// Location Number	Size	Description
// ----------------	----	-----------
// 000 - 001			-		reserved
// 002					250	buffer 1 (transmit)
// 003					250	buffer 2 (receive)
// 004 - 019			-		reserved
// 020 - 039			120	general purpose

type Memory [][]byte

var (
	SourceBufferLocation = ReceiveBufferLocation
	DestBufferLocation   = TransmitBufferLocation
)

func Run() {
	memory := InitMemory()
	Load(memory, "003=HELLO\n020=A\n021$X")

	// 020=A
	// 003=HELLO
	//
	// 021$
	//   X

	fmt.Println(string(memory[3]))
	fmt.Println(string(memory[20]))
	fmt.Println(string(memory[TransmitBufferLocation]))
}

func InitMemory() Memory {
	memory := make(Memory, MaxMemory)
	memory[TransmitBufferLocation] = make([]byte, 250)
	memory[ReceiveBufferLocation] = make([]byte, 250)

	for i := 4; i < MaxMemory; i++ {
		memory[i] = make([]byte, 120)
	}

	return memory
}

func Load(memory Memory, program string) {
	lines := strings.Split(program, "\n")
	memory = Assign(memory, lines[0])
	memory = Assign(memory, lines[1])

	tokens := strings.Split(lines[2], "$")
	// ???: Need to do something with location in tokens[0].
	memory = RunScript(memory, tokens[1])
}

func Assign(memory Memory, directive string) Memory {
	location := 0
	value := ""

	_, err := fmt.Sscanf(directive, "%d=%s", &location, &value)
	if err != nil {
		panic(err)
	}

	if len(value) > 120 {
		panic("value exceeds maximum size")
	}

	memory[location] = []byte(value)
	return memory
}

func RunScript(memory Memory, script string) Memory {
	if string(script) == "X" {
		memory[DestBufferLocation] = memory[SourceBufferLocation]
	}

	return memory
}
