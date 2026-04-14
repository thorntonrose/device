package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestV(t *testing.T) {
	memory := mem.New()
	memory.WriteAll(mem.Transmit, []byte("FOO"))
	memory.Reset(mem.Transmit)
	memory.WriteAll(mem.Receive, []byte("BAR"))
	memory.Reset(mem.Receive)

	v := NewV(memory)
	assert.Equal(t, "FOO", v.ReadAll([]string{}))
	assert.Equal(t, "BAR", v.ReadAll([]string{"2"}))
}
