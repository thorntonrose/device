package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

// Model 2?
//
// func TestIntFromSlot(t *testing.T) {
// 	c := New(mem.New())
// 	c.Memory.Slots[20] = []byte("123")
// 	c.Memory.Slots[21] = []byte("A")
//
// 	assert.Equal(t, 123, c.IntFromSlot("c", "#20"))
// 	assert.Panics(t, func() { c.IntFromSlot("c", "#A") })
// 	assert.Panics(t, func() { c.IntFromSlot("c", "#21") })
// }

func TestIntFromValue(t *testing.T) {
	c := New(mem.New())

	assert.Equal(t, 123, c.IntFromValue("c", "123"))
	assert.Panics(t, func() { c.IntFromValue("c", "A") })
}

func TestVarNum(t *testing.T) {
	c := New(mem.New())
	c.Memory.Variables[0] = 123

	assert.Equal(t, 123, c.VarNum("c", "#0"))
	assert.Panics(t, func() { c.VarNum("c", "#10") })
	assert.Panics(t, func() { c.VarNum("c", "#A") })
}

func TestIntFromParameter(t *testing.T) {
	c := New(mem.New())
	c.Memory.Variables[1] = 123

	assert.Equal(t, 123, c.IntFromParameter("c", "123"))
	assert.Equal(t, 123, c.IntFromParameter("c", "#1"))
}

func TestInt(t *testing.T) {
	c := New(mem.New())
	c.Memory.Variables[1] = 65

	assert.Equal(t, 1, c.Int("c", []string{}, 0, 1))
	assert.Equal(t, 1, c.Int("c", []string{"1"}, 0, 0))
	assert.Equal(t, 65, c.Int("c", []string{"#1"}, 0, 0))
	assert.Panics(t, func() { c.Int("c", []string{"#A"}, 0, 0) })
}

//-----------------------------------------------------------------------------

func TestCode(t *testing.T) {
	c := New(mem.New())

	assert.Equal(t, 1, c.Code("c", []string{"1"}, 0, 1, []int{1, 2}))
	assert.Equal(t, 1, c.Code("c", []string{"1"}, 0, 0, []int{1, 2}))
	assert.Panics(t, func() { c.Code("c", []string{"3"}, 0, 0, []int{1, 2}) })
}

func TestNonNegative(t *testing.T) {
	c := New(mem.New())

	assert.Equal(t, 0, c.NonNegative("c", []string{"0"}, 0, 1))
	assert.Equal(t, 1, c.NonNegative("c", []string{"1"}, 0, 0))
	assert.Panics(t, func() { c.NonNegative("c", []string{"-1"}, 0, 0) })
}

func TestPositive(t *testing.T) {
	c := New(mem.New())

	assert.Equal(t, 1, c.Positive("c", []string{"1"}, 0, 1))
	assert.Panics(t, func() { c.Positive("c", []string{"0"}, 0, 1) })
	assert.Panics(t, func() { c.Positive("c", []string{"-1"}, 0, 1) })
}

func TestRange(t *testing.T) {
	c := New(mem.New())

	assert.Equal(t, 5, c.Range("c", []string{"5"}, 0, 1, 1, 10))
	assert.Panics(t, func() { c.Range("c", []string{"0"}, 0, 1, 1, 10) })
	assert.Panics(t, func() { c.Range("c", []string{"11"}, 0, 1, 1, 10) })
}

func TestBytes(t *testing.T) {
	c := New(mem.New())

	assert.Equal(t, []byte{28}, c.Bytes("c", []string{}, 0, []byte{28}))
	assert.Equal(t, []byte{28}, c.Bytes("c", []string{""}, 0, []byte{28}))
	assert.Equal(t, []byte{65}, c.Bytes("c", []string{"65"}, 0, []byte{28}))
	assert.Equal(t, []byte("A1!"), c.Bytes("c", []string{"'A1!'"}, 0, []byte{28}))
	assert.Panics(t, func() { c.Bytes("c", []string{"256"}, 0, []byte{28}) })
	assert.Panics(t, func() { c.Bytes("c", []string{"A"}, 0, []byte{28}) })
}
