package mem

import (
	"fmt"
	"log"
	"slices"
	"strings"

	. "github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/iter"
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
// 003			250	buffer 2 (receive)
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

func NewBuffers(self *Memory) []*[]byte {
	buffers := make([]*[]byte, MaxBuffers+1)
	buffers[Transmit] = &self.Slots[Transmit+1]
	buffers[Receive] = &self.Slots[Receive+1]

	return buffers
}

//-----------------------------------------------------------------------------

func (self *Memory) Set(slotNum int, data []byte) {
	self.Slots[slotNum] = self.Slots[slotNum][:len(data)]
	copy(self.Slots[slotNum], data)
}

func (self *Memory) Load(data map[int][]byte) {
	log.Printf("Memory.Load")
	iter.EachEntry(data, func(slotNum int, value []byte) { self.Set(slotNum, value) })
}

//-----------------------------------------------------------------------------

func (self *Memory) Dump(slotNums ...int) string {
	lines := []string{
		fmt.Sprintf("SRC:  %d", self.Source),
		fmt.Sprintf("DEST: %d", self.Destination),
		fmt.Sprintf("PTRS: %v", self.Pointers[1:]),
		fmt.Sprintf("VARS: %v", self.Variables),
	}

	return strings.Join(append(lines, self.SlotLines(slotNums...)...), "\n")
}

func (self *Memory) SlotLines(slotNums ...int) (lines []string) {
	iter.EachWithIndex(self.Slots, func(data []byte, i int) {
		lines = append(lines, self.SlotLine(data, i, slotNums...)...)
	})

	return
}

func (self *Memory) SlotLine(data []byte, slotNum int, slotNums ...int) []string {
	includeLine := If(len(slotNums) == 0, len(data) > 0, slices.Contains(slotNums, slotNum))
	return If(includeLine, []string{fmt.Sprintf("%03d:  %s", slotNum, string(data))}, []string{})
}
