package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStarL(t *testing.T) {
	sl := NewStarL(nil, &MockScript{})

	AssertStarL(t, sl, []string{}, 0, 0)    // defaults
	AssertStarL(t, sl, []string{"2"}, 0, 2) // slot number

	assert.Panics(t, func() { sl.Run([]string{"A"}) })
}

func AssertStarL(t *testing.T, sl StarL, parameters []string, expectedSkip int, expectedSlotNum int) {
	assert.Equal(t, expectedSkip, sl.Run(parameters))
	assert.Equal(t, expectedSlotNum, sl.Script.(*MockScript).SlotNum)
}

//-----------------------------------------------------------------------------

type MockScript struct {
	SlotNum int
}

func (m *MockScript) Run(slotNum int) (skip int) {
	m.SlotNum = slotNum
	return
}
