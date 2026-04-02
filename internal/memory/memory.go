package memory

import "fmt"

const (
	MaxLocations    = 40
	MaxGeneralSize  = 120
	MaxReservedSize = 60
)

type Memory struct {
	Locations         []Location
	Buffers           []Buffer
	SourceBuffer      *Buffer
	DestinationBuffer *Buffer
	Variables         []byte
}

func New() Memory {
	memory := Memory{Locations: NewLocations()}
	memory.Buffers = NewBuffers(memory.Locations)
	memory.SourceBuffer = &memory.Buffers[ReceiveBufferNum-1]
	memory.DestinationBuffer = &memory.Buffers[TransmitBufferNum-1]

	return memory
}

//-----------------------------------------------------------------------------

func (m *Memory) Get(index int) Location {
	return m.Locations[index]
}

func (m *Memory) Set(index int, value []byte) {
	m.Locations[index].Set(value)
}

func (m *Memory) Dump() {
	fmt.Println("Location\tSize\tValue")

	for i, location := range m.Locations {
		fmt.Printf("%03d\t\t%d\t%s\n", i, cap(location), location.String())
	}
}
