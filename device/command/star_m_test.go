package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/device/mem"
)

func TestStarM(t *testing.T) {
	sm := NewStarM(nil)
	assert.Equal(t, mem.MaxBufferSize+1, sm.Run([]string{}))
}
