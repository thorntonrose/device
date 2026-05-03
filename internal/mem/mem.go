package mem

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/thorntonrose/device/internal/etc"
)

const (
	MaxSlots     = 40
	MaxBuffers   = 2
	MaxVariables = 10

	MaxBufferSize   = 250
	MaxGeneralSize  = 120
	MaxReservedSize = 60

	Transmit = 1
	Receive  = 2
)

// Slot #		Size	Description
// --------		----	-----------
// 000 - 001	-		reserved
// 002			250	buffer 1 (transmit)
// 003			250	buffer 2 (selfeive)
// 004 - 019	-		reserved
// 020 - 039	120	general purpose
type Memory struct {
	Slots       [][]byte
	Buffers     []*[]byte
	Pointers    []int
	Variables   []int
	Source      int
	Destination int
}

func New() *Memory {
	m := &Memory{Slots: NewSlots(), Pointers: make([]int, MaxBuffers+1), Variables: make([]int, MaxVariables),
		Source: Receive, Destination: Transmit}
	m.Buffers = NewBuffers(m)

	return m
}

func NewSlots() [][]byte {
	slots := make([][]byte, MaxSlots)
	AddBlock(&slots, 0, 1, MaxReservedSize)
	AddBlock(&slots, 2, 3, MaxBufferSize)
	AddBlock(&slots, 4, 19, MaxReservedSize)
	AddBlock(&slots, 20, 39, MaxGeneralSize)

	return slots
}

func AddBlock(slots *[][]byte, start int, end int, size int) {
	for i := start; i <= end; i++ {
		AddSlot(slots, i, size)
	}
}

func AddSlot(slots *[][]byte, i int, size int) {
	(*slots)[i] = make([]byte, 0, size)

	// ???: Set slot to a generated word randomly.
	// if size == MaxReservedSize {
	// 	(*slots)[i] = append((*slots)[i], fmt.Sprintf("%016x", rand.Intn(math.MaxInt64-1)+1)...)
	// }
}

func NewBuffers(m *Memory) []*[]byte {
	buffers := make([]*[]byte, MaxBuffers+1)
	buffers[Transmit] = &m.Slots[Transmit+1]
	buffers[Receive] = &m.Slots[Receive+1]

	return buffers
}

//-----------------------------------------------------------------------------

func (m *Memory) Set(slotNum int, data []byte) {
	m.Slots[slotNum] = m.Slots[slotNum][:len(data)]
	copy(m.Slots[slotNum], data)
}

// ???: Needed?
// func (m *Memory) Append(slotNum int, data []byte) {
// 	size := len(data)
// 	m.Slots[slotNum] = m.Slots[slotNum][:len(m.Slots[slotNum])+size]
// 	copy(m.Slots[slotNum][len(m.Slots[slotNum])-size:], data)
// }

func (m *Memory) Load(data map[int][]byte) {
	log.Printf("Memory.Load: data: %v\n", data)
	etc.EachEntry(data, func(slotNum int, value []byte) { m.Set(slotNum, value) })
}

//-----------------------------------------------------------------------------

func (m *Memory) Dump(slotNums ...int) string {
	return strings.Join(append([]string{m.BufferLine()}, m.SlotLines(slotNums...)...), "\n")
}

func (m *Memory) SlotLines(slotNums ...int) (lines []string) {
	etc.EachWithIndex(m.Slots, func(data []byte, i int) { lines = append(lines, m.SlotLine(data, i, slotNums...)...) })
	return
}

func (m *Memory) SlotLine(data []byte, slotNum int, slotNums ...int) []string {
	includeLine := etc.If(len(slotNums) == 0, len(data) > 0, slices.Contains(slotNums, slotNum))
	return etc.If(includeLine, []string{fmt.Sprintf("%03d (%03d): %s", slotNum, cap(data), string(data))}, []string{})
}

func (m *Memory) BufferLine() string {
	return fmt.Sprintf("S: %d, D: %d, P: %v", m.Source, m.Destination, m.Pointers)
}
