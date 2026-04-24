package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestInt(t *testing.T) {
	c := NewCommand(mem.New())
	assert.Equal(t, 1, c.Int("c", []string{}, 0, 1))
	assert.Equal(t, 123, c.Int("c", []string{"123"}, 0, 0))
	assert.Panics(t, func() { c.Int("c", []string{"A"}, 0, 0) })
}
