package bufs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestO(t *testing.T) {
	o := NewO(mem.New())
	o.Memory.Set(mem.Receive+1, []byte("FOO"))

	for i, c := range []string{"", "1"} {
		o.Run([]string{c})
		assert.Equal(t, i+1, o.Memory.Pointers[o.Memory.Source])
	}
}
