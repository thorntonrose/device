package command

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/thorntonrose/device/device/etc"
	"github.com/thorntonrose/device/device/mem"
	_ "github.com/thorntonrose/device/device/testing"
)

func TestPlusI_Transmit(t *testing.T) {
	pi := NewPlusI(mem.New())
	*pi.Memory.Buffers[mem.Transmit] = []byte("FOO")

	AssertTransmit(t, pi, "", "FOO")
	AssertTransmit(t, pi, "1", "FOO\n")
	assert.Panics(t, func() { pi.Run([]string{"6"}) })
}

func AssertTransmit(t *testing.T, pi PlusI, a string, expected string) {
	reader, writer, restore := StdoutPipe()
	defer restore()

	pi.Run([]string{a})
	writer.Close()
	assert.Equal(t, expected, string(Must(io.ReadAll(reader))))
}

func StdoutPipe() (*os.File, *os.File, func()) {
	origStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	return reader, writer, func() { os.Stdout = origStdout; writer.Close() }
}

//-----------------------------------------------------------------------------

func TestPlusI_Receive(t *testing.T) {
	pi := NewPlusI(mem.New())

	writer, restore := StdinPipe()
	defer restore()
	writer.Write([]byte("FOO"))
	writer.Close()

	skip := pi.Run([]string{"5", "0", "1", "3"})
	assert.Equal(t, []byte("FOO"), *pi.Memory.Buffers[mem.Receive])
	assert.Equal(t, 1, skip)

	assert.Panics(t, func() { pi.Run([]string{"", "", "", fmt.Sprintf("%d", mem.MaxBufferSize+1)}) })
}

func StdinPipe() (*os.File, func()) {
	origStdin := os.Stdin
	reader, writer, _ := os.Pipe()
	os.Stdin = reader

	return writer, func() { os.Stdin = origStdin; reader.Close(); writer.Close() }
}
