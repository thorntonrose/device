package mem

import (
	"fmt"
	"strings"

	"github.com/thorntonrose/device/internal/etc"
)

const (
	MaxBuffers   = 2
	MaxSlots     = 40
	MaxVariables = 10

	MaxBufferSize   = 250
	MaxGeneralSize  = 120
	MaxReservedSize = 60

	Transmit = 1
	Receive  = 2
)

// Slot			Size	Description
// ---------	----	-----------
// 000 - 001	-		reserved
// 002			250	buffer 1 (transmit)
// 003			250	buffer 2 (receive)
// 004 - 019	-		reserved
// 020 - 039	120	general purpose

type Memory struct {
	Slots        [][]byte
	Variables    []byte
	ReadPointers map[int]int
	Source       int
	Destination  int
}

func New() Memory {
	return Memory{
		Slots:        NewSlots(),
		Variables:    make([]byte, MaxVariables),
		ReadPointers: map[int]int{Receive: 0, Transmit: 0},
		Source:       Receive,
		Destination:  Transmit,
	}
}

func NewSlots() [][]byte {
	// ???: Set reserved slots to special value?
	slots := make([][]byte, MaxSlots)
	AddBlock(&slots, 0, 1, MaxReservedSize)
	AddBlock(&slots, 2, 3, MaxBufferSize)
	AddBlock(&slots, 4, 19, MaxReservedSize)
	AddBlock(&slots, 20, 39, MaxGeneralSize)

	return slots
}

func AddBlock(slots *[][]byte, start int, end int, size int) {
	for i := start; i <= end; i++ {
		(*slots)[i] = make([]byte, 0, size)
	}
}

func Slot(buf int) int {
	return buf + 1
}

//-----------------------------------------------------------------------------

func (m Memory) Get(slot int) []byte {
	return m.Slots[slot]
}

func (m *Memory) Set(slot int, data []byte) {
	m.Slots[slot] = m.Slots[slot][:len(data)]
	copy(m.Slots[slot], data)
}

func (m *Memory) Append(slot int, data []byte) {
	size := len(data)
	m.Slots[slot] = m.Slots[slot][:len(m.Slots[slot])+size]
	copy(m.Slots[slot][len(m.Slots[slot])-size:], data)
}

func (m Memory) Load(data map[int][]byte) {
	etc.EachEntry(data, func(slot int, value []byte) { m.Set(slot, value) })
}

func (m Memory) Dump() string {
	lines := m.DumpSlots()
	pointers := strings.TrimPrefix(fmt.Sprintf("%v", m.ReadPointers), "map")
	lines = append(lines, fmt.Sprintf("S: %d, D: %d, P: %v", m.Source, m.Destination, pointers))

	return strings.Join(lines, "\n")
}

func (m Memory) DumpSlots() (lines []string) {
	etc.EachWithIndex(m.Slots, func(slot []byte, i int) {
		lines = append(lines, etc.If(len(slot) == 0, []string{}, []string{m.DumpSlot(i)})...)
	})

	return lines
}

func (m Memory) DumpSlot(slot int) string {
	return fmt.Sprintf("%03d (%03d): %s", slot, cap(m.Slots[slot]), string(m.Slots[slot]))
}

//-----------------------------------------------------------------------------

func (m *Memory) ReadAll(buf int, maxCount int, stop byte) (data []byte) {
	for count := 0; m.HasNext(buf, count, maxCount); count++ {
		b := m.Read(buf)

		if stop > 0 && b == stop {
			break
		}

		data = append(data, b)
	}

	return data
}

func (m *Memory) HasNext(buf int, count int, maxCount int) bool {
	return (maxCount == 0 || count < maxCount) && m.ReadPointers[buf] < len(m.Slots[Slot(buf)])
}

func (m *Memory) Read(buf int) byte {
	data := m.Slots[Slot(buf)][m.ReadPointers[buf]]
	m.ReadPointers[buf]++

	return data
}

func (m *Memory) WriteAll(buf int, data []byte) {
	m.Append(Slot(buf), data)
	m.ReadPointers[buf] += len(data)
}

func (m *Memory) Clear(buf int) {
	slot := Slot(buf)
	m.Slots[slot] = m.Slots[slot][:0]
	m.Reset(buf)
}

func (m *Memory) Reset(buf int) {
	m.ReadPointers[buf-1] = 0
}
