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
	MaxBuffers      = 5
	MaxBufferSize   = 250
	MaxGeneralSize  = 120
	MaxReservedSize = 60
	MaxSlots        = 1000
	MaxVariables    = 10

	Transmit = 1
	Receive  = 2
)

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
	AddSlots(&slots, 0, 1, MaxReservedSize)
	AddSlots(&slots, 2, 6, MaxBufferSize)
	AddSlots(&slots, 7, 19, MaxReservedSize)
	AddSlots(&slots, 20, 99, MaxGeneralSize)
	AddSlots(&slots, 100, 112, MaxReservedSize)
	AddSlots(&slots, 113, 949, MaxGeneralSize)
	AddSlots(&slots, 950, 999, MaxReservedSize)

	return slots
}

func AddSlots(slots *[][]byte, start int, end int, size int) {
	for i := start; i <= end; i++ {
		(*slots)[i] = make([]byte, 0, size)
	}
}

func NewBuffers(self *Memory) []*[]byte {
	buffers := make([]*[]byte, MaxBuffers+1)
	buffers[Transmit] = &self.Slots[Transmit+1]
	buffers[Receive] = &self.Slots[Receive+1]

	return buffers
}

//-----------------------------------------------------------------------------

func (self *Memory) Load(data map[int][]byte) {
	log.Println("Load")
	iter.EachEntry(data, func(slotNum int, value []byte) { self.Set(slotNum, value) })
}

func (self *Memory) Set(slotNum int, data []byte) {
	self.Slots[slotNum] = self.Slots[slotNum][:len(data)]
	copy(self.Slots[slotNum], data)
}

//-----------------------------------------------------------------------------

func (self *Memory) Dump(slotNums ...int) string {
	log.Println("Dump")

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
