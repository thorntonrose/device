package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestStarN(t *testing.T) {
	memory := mem.New()
	starN := NewStarN(memory)
	starN.Memory.Variables[0] = 1

	starN.Run([]string{})
	assert.Equal(t, 0, memory.Variables[0])

	starN.Run([]string{"", "1"})
	assert.Equal(t, 1, memory.Variables[0])

	starN.Run([]string{"#1", "1"})
	assert.Equal(t, 1, memory.Variables[1])

	assert.Panics(t, func() { starN.Run([]string{"1"}) })
	assert.Panics(t, func() { starN.Run([]string{"#10"}) })
}
