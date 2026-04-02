package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BufferSuite struct {
	suite.Suite
}

func TestBufferSuite(t *testing.T) {
	suite.Run(t, new(BufferSuite))
}

//-----------------------------------------------------------------------------

func (s *BufferSuite) TestNewBuffers() {
	memory := New()
	buffers := memory.Buffers

	assert.Len(s.T(), buffers, MaxBuffers)
	assert.NotNil(s.T(), buffers[TransmitBufferNum-1])
	assert.NotNil(s.T(), buffers[ReceiveBufferNum-1])
}

//-----------------------------------------------------------------------------

func (s *BufferSuite) TestReset() {
	buffer := &Buffer{ExtractionPointer: 5}
	buffer.Reset()
	assert.Equal(s.T(), 0, buffer.ExtractionPointer)
}

func (s *BufferSuite) TestWrite() {
	location := NewLocation(MaxBufferSize)
	buffer := NewBuffer(&location)

	buffer.Write(byte("F"[0]))
	assert.Equal(s.T(), "F", location.String())
	assert.Equal(s.T(), 1, buffer.ExtractionPointer)
}

func (s *BufferSuite) TestRead() {
	location := NewLocation(MaxBufferSize)

	buffer := NewBuffer(&location)
	buffer.Write(byte("F"[0]))
	buffer.Reset()

	data := buffer.Read()
	assert.Equal(s.T(), "F", string(data))
	assert.Equal(s.T(), 1, buffer.ExtractionPointer)
}
