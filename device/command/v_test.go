package command

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	// . "github.com/thorntonrose/device/device/etc"
	"github.com/thorntonrose/device/device/mem"
	_ "github.com/thorntonrose/device/device/testing"
)

func TestV(t *testing.T) {
	v := NewV(mem.New())
	*v.Memory.Buffers[mem.Transmit] = []byte("FOO")
	*v.Memory.Buffers[mem.Receive] = []byte("BAR")

	AssertV(t, v, []string{""}, "FOO")  // defaults
	AssertV(t, v, []string{"2"}, "BAR") // receive buffer

	assert.Panics(t, func() { v.Run([]string{fmt.Sprintf("%d", mem.MaxBuffers+1)}) })
}

func AssertV(t *testing.T, v V, parameters []string, expected string) {
	buf := &strings.Builder{}
	Display = buf

	v.Run(parameters)
	assert.Equal(t, expected+"\n", buf.String())
}
