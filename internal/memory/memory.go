package memory

import "fmt"

const (
	MaxLocations   = 40
	MaxGeneralSize = 120
)

type Memory struct {
	// ???: Make []byte a type with reserved flag?
	Locations         [][]byte
	Buffers           []Buffer
	SourceBuffer      *Buffer
	DestinationBuffer *Buffer
}

func New() Memory {
	memory := Memory{Locations: NewLocations(), Buffers: NewBuffers()}
	// ???: Do this in buffers.go?
	memory.Buffers[TransmitBufferNum-1].Location = &memory.Locations[TransmitBufferNum+1]
	memory.Buffers[ReceiveBufferNum-1].Location = &memory.Locations[ReceiveBufferNum+1]
	memory.SourceBuffer = &memory.Buffers[ReceiveBufferNum-1]
	memory.DestinationBuffer = &memory.Buffers[TransmitBufferNum-1]

	return memory
}

func NewLocations() [][]byte {
	locations := make([][]byte, MaxLocations)
	InitLocations(&locations, 2, MaxBuffers, MaxBufferSize)
	InitLocations(&locations, 20, MaxLocations-20, MaxGeneralSize)
	// ???: Set reserved locations to special value?

	return locations
}

func InitLocations(location *[][]byte, start int, count int, size int) {
	for i := start; i < start+count; i++ {
		(*location)[i] = make([]byte, 0, size)
	}
}

//-----------------------------------------------------------------------------

func (m *Memory) Get(locNum int) []byte {
	m.CheckBounds(locNum)
	return m.Locations[locNum]
}

func (m *Memory) Set(locNum int, value []byte) {
	m.CheckBounds(locNum)
	m.CheckReserved(locNum)

	loc := &m.Locations[locNum]
	m.CheckSize(*loc, value)
	*loc = (*loc)[:len(value)]
	copy(*loc, value)
}

func (m *Memory) CheckBounds(locNum int) {
	if locNum < 0 || locNum >= MaxLocations {
		panic("location number out of bounds")
	}
}

func (m *Memory) CheckSize(location []byte, value []byte) {
	if len(value) > cap(location) {
		panic("value exceeds maximum size")
	}
}

func (m *Memory) CheckReserved(locNum int) {
	if m.IsReserved(locNum) {
		panic("location is reserved")
	}
}

func (m *Memory) IsReserved(locNum int) bool {
	return m.Locations[locNum] == nil
}

func (m *Memory) Dump() {
	fmt.Println("Location\tSize\tValue")

	for i, loc := range m.Locations {
		fmt.Printf("%03d\t\t%d\t%s\n", i, cap(loc), string(loc))
	}
}
