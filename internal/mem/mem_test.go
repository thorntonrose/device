package mem

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/config"
)

func TestMain(m *testing.M) {
	defer config.InitLog()()
	m.Run()
}

//-----------------------------------------------------------------------------

func TestNew(t *testing.T) {
	memory := New()
	assert.Len(t, memory.Slots, MaxSlots)
	// common to all device models
	assertBlock(t, memory.Slots, 0, 1, MaxReservedSize)
	assertBlock(t, memory.Slots, 2, 3, MaxBufferSize)
	assertBlock(t, memory.Slots, 20, 99, MaxGeneralSize)
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
	assert.Equal(t, data, memory.Slots[20])
}

func TestSet(t *testing.T) {
	data := []byte("FOO")
	memory := New()

	memory.Set(20, data)
	assert.Equal(t, data, memory.Slots[20])
}

func TestLoad(t *testing.T) {
	memory := New()
	data := map[int][]byte{20: []byte("FOO"), 21: []byte("BAR")}

	memory.Load(data)
	assert.Equal(t, data[20], memory.Slots[20])
	assert.Equal(t, data[21], memory.Slots[21])
}

func TestDump(t *testing.T) {
	memory := New()
	memory.Set(2, []byte("A"))
	memory.Set(3, []byte("B"))
	memory.Set(20, []byte("C"))

	lines := []string{
		"SRC:  2",
		"DEST: 1",
		"PTRS: [0 0 0 0 0]",
		"VARS: [0 0 0 0 0 0 0 0 0 0]",
		"002:  A",
		"003:  B",
		"020:  C",
	}

	assert.Equal(t, strings.Join(lines, "\n"), memory.Dump(2, 3, 20))
}
