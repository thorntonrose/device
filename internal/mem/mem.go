package mem

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"strings"
	"time"

	"github.com/thorntonrose/device/internal/etc"
)

const (
	MaxBuffers      = 2
	MaxSlots        = 40
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
	Slots       [][]byte
	Pointers    []int
	Source      int
	Destination int
}

func New() *Memory {
	return &Memory{Slots: NewSlots(), Pointers: make([]int, MaxBuffers+1), Source: Receive, Destination: Transmit}
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
		AddBufSlot(slots, i, size)
	}
}

func AddBufSlot(slots *[][]byte, i int, size int) {
	(*slots)[i] = make([]byte, 0, size)

	if size == MaxReservedSize {
		(*slots)[i] = append((*slots)[i], RandomValue()...)
	}
}

func RandomValue() []byte {
	md5sum := md5.Sum([]byte(time.Now().String()))
	s := hex.EncodeToString(md5sum[:])

	return []byte(strings.ToUpper(s)[:rand.Intn(len(s)-1)+1])
}

func BufSlot(buf int) int {
	return buf + 1
}

//-----------------------------------------------------------------------------

func (m *Memory) Get(slot int) []byte {
	return m.Slots[slot]
}

func (m *Memory) Set(slot int, data []byte) {
	m.Slots[slot] = m.Slots[slot][:len(data)]
	copy(m.Slots[slot], data)
}

// ???: Needed?
// func (m *Memory) Append(slot int, data []byte) {
// 	size := len(data)
// 	m.Slots[slot] = m.Slots[slot][:len(m.Slots[slot])+size]
// 	copy(m.Slots[slot][len(m.Slots[slot])-size:], data)
// }

func (m *Memory) Load(data map[int][]byte) {
	log.Printf("Memory.Load: data: %v\n", data)
	etc.EachEntry(data, func(slot int, value []byte) { m.Set(slot, value) })
}

//-----------------------------------------------------------------------------

func (m *Memory) Dump(slots ...int) string {
	return strings.Join(append([]string{m.BufferLine()}, m.SlotLines(slots...)...), "\n")
}

func (m *Memory) SlotLines(slots ...int) (lines []string) {
	etc.EachWithIndex(m.Slots, func(data []byte, i int) { lines = append(lines, m.SlotLine(data, i, slots...)...) })
	return lines
}

func (m *Memory) SlotLine(data []byte, slot int, slots ...int) []string {
	if len(slots) == 0 || slices.Contains(slots, slot) {
		return []string{fmt.Sprintf("%03d (%03d): %s", slot, cap(data), string(data))}
	}

	return []string{}
}

func (m *Memory) BufferLine() string {
	return fmt.Sprintf("S: %d, D: %d, P: %v", m.Source, m.Destination, m.Pointers)
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
	return (maxCount == 0 || count < maxCount) && m.Pointers[buf] < len(m.Slots[BufSlot(buf)])
}

func (m *Memory) Read(buf int) byte {
	data := m.Slots[BufSlot(buf)][m.Pointers[buf]]
	m.Pointers[buf]++

	return data
}

func (m *Memory) WriteAll(buf int, data []byte) {
	etc.Each(data, func(b byte) { m.Write(buf, b) })
}

func (m *Memory) Write(buf int, data byte) {
	m.Slots[BufSlot(buf)] = append(m.Slots[BufSlot(buf)], data)
	m.Pointers[buf]++
}

func (m *Memory) Clear(buf int) {
	slot := BufSlot(buf)
	m.Slots[slot] = m.Slots[slot][:0]
	m.Reset(buf)
}

func (m *Memory) Reset(buf int) {
	m.Pointers[buf] = 0
}
