package etc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	assert.NotPanics(t, func() { Check(nil) })
	assert.Panics(t, func() { Check(assert.AnError) })
}

func TestEach(t *testing.T) {
	result := []int{}
	Each([]int{1, 2, 3}, func(i int) { result = append(result, i) })
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestEachWithIndex(t *testing.T) {
	result := [2]string{}
	EachWithIndex([]string{"a", "b"}, func(s string, i int) { result[i] = fmt.Sprintf("%s%d", s, i) })
	assert.Equal(t, [2]string{"a0", "b1"}, result)
}

func TestEachEntry(t *testing.T) {
	result := []string{}
	EachEntry(map[string]string{"a": "1", "b": "2"}, func(k string, v string) { result = append(result, k+v) })
	assert.Contains(t, result, "a1")
	assert.Contains(t, result, "b2")
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
