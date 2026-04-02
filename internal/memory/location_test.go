package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LocationSuite struct {
	suite.Suite
}

func TestLocationSuite(t *testing.T) {
	suite.Run(t, new(LocationSuite))
}

//-----------------------------------------------------------------------------

func (s *LocationSuite) TestNewLocations() {
	locations := NewLocations()
	assert.Len(s.T(), locations, MaxLocations)
	assertCapacity(s.T(), locations, 0, 1, MaxReservedSize)
	assertCapacity(s.T(), locations, 2, 3, MaxBufferSize)
	assertCapacity(s.T(), locations, 4, 19, MaxReservedSize)
	assertCapacity(s.T(), locations, 20, 39, MaxGeneralSize)
}

func assertCapacity(t *testing.T, locations []Location, start int, end int, expected int) {
	for i := start; i <= end; i++ {
		assert.Equal(t, expected, cap(locations[i]))
	}
}

//-----------------------------------------------------------------------------

func (s *LocationSuite) TestAppend() {
	location := NewLocation(10)
	location.Append('A')
	assert.Equal(s.T(), "A", location.String())
}

func (s *LocationSuite) TestSet() {
	location := NewLocation(10)
	data := []byte("FOO")
	location.Set(data)

	assert.Equal(s.T(), "FOO", location.String())
}
