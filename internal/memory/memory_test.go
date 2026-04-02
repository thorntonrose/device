package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MemorySuite struct {
	suite.Suite
}

func TestMemorySuite(t *testing.T) {
	suite.Run(t, new(MemorySuite))
}

//-----------------------------------------------------------------------------

func (s *MemorySuite) TestNew() {
	memory := New()
	assert.Len(s.T(), memory.Locations, MaxLocations)
	assert.Len(s.T(), memory.Buffers, MaxBuffers)
	assert.NotNil(s.T(), memory.SourceBuffer)
	assert.NotNil(s.T(), memory.DestinationBuffer)
}

func (s *MemorySuite) TestSet() {
	data := []byte("FOO")
	memory := New()

	memory.Set(20, data)
	assert.Equal(s.T(), string(data), memory.Locations[20].String())
}

func (s *MemorySuite) TestGet() {
	data := []byte("FOO")
	memory := New()

	memory.Set(20, data)
	assert.Equal(s.T(), string(data), memory.Get(20).String())
}
