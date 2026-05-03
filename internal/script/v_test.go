package script

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
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
	reader, writer, restore := StderrPipe()
	defer restore()

	v.Run(parameters)
	writer.Close()
	assert.Equal(t, expected+"\n", string(etc.Must(io.ReadAll(reader))))
}
