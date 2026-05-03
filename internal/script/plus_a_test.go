package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestPlusA(t *testing.T) {
	pa := NewPlusA(mem.New())
	pa.Memory.Variables[1] = 1

	assert.Equal(t, 0, pa.Run([]string{}))                    // defaults
	assert.Equal(t, 1, pa.Run([]string{"#1", "0", "1", "1"})) // equal
	assert.Equal(t, 1, pa.Run([]string{"#1", "1", "0", "1"})) // greater, true
	assert.Equal(t, 0, pa.Run([]string{"#1", "1", "1", "1"})) // greater, false
	assert.Equal(t, 1, pa.Run([]string{"#1", "2", "2", "1"})) // less, true
	assert.Equal(t, 0, pa.Run([]string{"#1", "2", "1", "1"})) // less, false

	assert.Panics(t, func() { pa.Run([]string{"#0", "3"}) }) // invalid
}
