package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
	_ "github.com/thorntonrose/device/internal/testing"
)

func TestStarN(t *testing.T) {
	sn := NewStarN(mem.New())
	sn.Memory.Variables[0] = 1

	sn.Run([]string{})
	assert.Equal(t, 0, sn.Memory.Variables[0])

	sn.Run([]string{"#1", "2"})
	assert.Equal(t, 2, sn.Memory.Variables[1])
}

func TestStarN_FromVariable(t *testing.T) {
	sn := NewStarN(mem.New())
	sn.Memory.Variables[2] = 2

	sn.Run([]string{"#1", "#2"})
	assert.Equal(t, 2, sn.Memory.Variables[1])
}
