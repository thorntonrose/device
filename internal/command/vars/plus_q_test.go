package vars

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestPlusQ(t *testing.T) {
	pq := NewPlusQ(mem.New())
	pq.Memory.Variables[0] = 2748
	pq.Memory.Variables[1] = 86
	pq.Memory.Variables[2] = 22617

	AssertPlusQ(t, pq, []string{}, []byte("2748"))        // string (2748 -> “2748”)
	AssertPlusQ(t, pq, []string{"#1", "1"}, []byte("V"))  // ASCII (86 -> “V”)
	AssertPlusQ(t, pq, []string{"#2", "1"}, []byte("XY")) // ASCII (22617 -> “XY”)

	assert.Panics(t, func() { pq.Run([]string{"#10"}) })
	assert.Panics(t, func() { pq.Run([]string{"", "2"}) })
}

func AssertPlusQ(t *testing.T, pq PlusQ, parameters []string, expected []byte) {
	pq.Memory.Clear(mem.Transmit)
	pq.Run(parameters)
	assert.Equal(t, expected, *pq.Memory.Buffers[mem.Transmit])
}
