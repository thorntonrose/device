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

func TestP(t *testing.T) {
	p := NewP(mem.New())
	p.Memory.Set(0, []byte("FOO"))
	p.Memory.Set(20, []byte("BAR"))

	AssertP(t, p, []string{}, "000:FOO")
	AssertP(t, p, []string{"20"}, "020:BAR")

	assert.Panics(t, func() { p.Run([]string{fmt.Sprintf("%d", mem.MaxSlots)}) })
}

func AssertP(t *testing.T, p P, parameters []string, expected string) {
	reader, writer, restore := StderrPipe()
	defer restore()

	p.Run(parameters)
	writer.Close()
	assert.Equal(t, expected+"\n", string(Must(io.ReadAll(reader))))
}

func StderrPipe() (*os.File, *os.File, func()) {
	origStderr := os.Stderr
	reader, writer, _ := os.Pipe()
	os.Stderr = writer

	return reader, writer, func() { os.Stderr = origStderr; writer.Close() }
}
