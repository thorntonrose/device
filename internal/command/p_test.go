package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestP(t *testing.T) {
	memory := mem.New()
	memory.Set(20, []byte("HELLO"))

	p := NewP(memory)
	assert.Equal(t, string(memory.Get(0)), p.Get([]string{}))
	assert.Equal(t, "HELLO", p.Get([]string{"20"}))

	p.Run([]string{"20"})
}
