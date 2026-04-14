package mem

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	memory := New()
	assert.Len(t, memory.Slots, MaxSlots)
	assert.Len(t, memory.Slots[0], MaxBuffers+1)
	assertBlock(t, memory.Slots, 1, 1, MaxReservedSize)
	assertBlock(t, memory.Slots, 2, 3, MaxBufferSize)
	assertBlock(t, memory.Slots, 4, 19, MaxReservedSize)
	assertBlock(t, memory.Slots, 20, 39, MaxGeneralSize)
}

func assertBlock(t *testing.T, slots [][]byte, start int, end int, size int) {
	for i := start; i <= end; i++ {
		assert.True(t, (size == MaxReservedSize && len(slots[i]) > 0) || len(slots[i]) == 0, "slot %d", i)
		assert.Equal(t, cap(slots[i]), size, "slot %d", i)
	}
}

//-----------------------------------------------------------------------------

func TestGet(t *testing.T) {
	data := []byte("FOO")
	memory := New()

	memory.Set(20, data)
	assert.Equal(t, data, memory.Get(20))
}

func TestSet(t *testing.T) {
	data := []byte("FOO")
	memory := New()

	memory.Set(20, data)
	assert.Equal(t, data, memory.Slots[20])
}

// func TestAppend(t *testing.T) {
// 	memory := New()

// 	memory.Append(20, []byte{'A'})
// 	assert.Equal(t, []byte{'A'}, memory.Slots[20])
// }

func TestLoad(t *testing.T) {
	memory := New()
	data := map[int][]byte{20: []byte("FOO"), 21: []byte("BAR")}

	memory.Load(data)
	assert.Equal(t, data[20], memory.Get(20))
	assert.Equal(t, data[21], memory.Get(21))
}

func TestDump(t *testing.T) {
	memory := New()
	memory.Set(2, []byte("A"))
	memory.Set(3, []byte("B"))
	memory.Set(20, []byte("C"))

	lines := []string{"S: 2, D: 1, P: [0 0 0]", "002 (250): A", "003 (250): B", "020 (120): C"}
	assert.Equal(t, strings.Join(lines, "\n"), memory.Dump(2, 3, 20))
}

//-----------------------------------------------------------------------------

func TestRead(t *testing.T) {
	memory := New()
	memory.Set(BufSlot(Receive), []byte{'A'})

	assert.Equal(t, byte('A'), memory.Read(Receive))
	assert.Equal(t, 1, memory.Pointers[Receive])
}

func TestReadAll(t *testing.T) {
	memory := New()
	memory.Set(BufSlot(Receive), []byte{'A', 'B', 'C'})

	assert.Equal(t, []byte{'A', 'B', 'C'}, memory.ReadAll(Receive, 0, 0))
}

func TestReadAll_Max(t *testing.T) {
	memory := New()
	memory.Set(BufSlot(Receive), []byte{'A', 'B', 'C'})

	assert.Equal(t, []byte{'A', 'B'}, memory.ReadAll(Receive, 2, 0))
}

func TestReadAll_Stop(t *testing.T) {
	memory := New()
	memory.Set(BufSlot(Receive), []byte{'A', 28, 'B'})

	assert.Equal(t, []byte{'A'}, memory.ReadAll(Receive, 0, 28))
}

func TestWriteAll(t *testing.T) {
	memory := New()

	memory.WriteAll(Transmit, []byte{'A'})
	assert.Equal(t, []byte{'A'}, memory.Slots[BufSlot(Transmit)])
}

func TestClear(t *testing.T) {
	memory := New()
	memory.WriteAll(Transmit, []byte{'A'})

	memory.Clear(Transmit)
	assert.Equal(t, []byte{}, memory.Slots[BufSlot(Transmit)])
}
