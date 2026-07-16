package etc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	assert.NotPanics(t, func() { Check(nil) })
	assert.Panics(t, func() { Check(assert.AnError) })
}

func TestIf(t *testing.T) {
	assert.Equal(t, "true", If(true, "true", "false"))
	assert.Equal(t, "false", If(false, "true", "false"))
}

func TestMust(t *testing.T) {
	assert.Equal(t, 1, Must(1, nil))
	assert.Panics(t, func() { Must(0, assert.AnError) })
}

func TestRecover(t *testing.T) {
	defer Recover(func(e error) { assert.NotNil(t, e) })
	panic("test")
}

func TestValue(t *testing.T) {
	assert.Equal(t, "value", Value("value", "default"))
	assert.Equal(t, "default", Value("", "default"))
}
