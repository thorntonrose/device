package mem

import (
	"fmt"
	"strings"
)

const (
	MaxLocations    = 40
	MaxBufferSize   = 250
	MaxGeneralSize  = 120
	MaxReservedSize = 60

	Transmit = 2
	Receive  = 3
)

// NewLocations creates a new set of memory locations.
//
// Location Number	Size	Description
// ---------------	----	-----------
// 000 - 001			-		reserved
// 002					250	buffer 1 (transmit)
// 003					250	buffer 2 (receive)
// 004 - 019			-		reserved
// 020 - 039			120	general purpose

type Memory [][]byte

func New() Memory {
	// ???: Set reserved locations to special value?
	memory := make(Memory, MaxLocations)
	memory = AddLocations(memory, 0, 1, MaxReservedSize)
	memory = AddLocations(memory, 2, 3, MaxBufferSize)
	memory = AddLocations(memory, 4, 19, MaxReservedSize)
	memory = AddLocations(memory, 20, 39, MaxGeneralSize)

	return memory
}

func AddLocations(memory Memory, start int, end int, size int) Memory {
	for i := start; i <= end; i++ {
		memory[i] = make([]byte, 0, size)
	}

	return memory
}

//-----------------------------------------------------------------------------

func (m Memory) Set(index int, data []byte) {
	m[index] = m[index][:len(data)]
	copy(m[index], data)
}

func (m Memory) Append(index int, data byte) {
	m[index] = m[index][:len(m[index])+1]
	m[index][len(m[index])-1] = data
}

func (m Memory) Dump(locations ...int) string {
	lines := []string{}

	for _, loc := range locations {
		lines = append(lines, fmt.Sprintf("%03d (%03d): %s", loc, cap(m[loc]), string(m[loc])))
	}

	return strings.Join(lines, "\n")
}
