package memory

const (
	MaxBuffers    = 2
	MaxBufferSize = 250

	TransmitBufferNum = 1
	ReceiveBufferNum  = 2
)

type Buffer struct {
	Location          *Location
	ExtractionPointer int
}

func NewBuffers(locations []Location) []Buffer {
	buffers := make([]Buffer, MaxBuffers)
	buffers[TransmitBufferNum-1] = *NewBuffer(&locations[TransmitBufferNum-1])
	buffers[ReceiveBufferNum-1] = *NewBuffer(&locations[ReceiveBufferNum-1])

	return buffers
}

func NewBuffer(location *Location) *Buffer {
	return &Buffer{Location: location}
}

//-----------------------------------------------------------------------------

func (b *Buffer) Read() byte {
	data := (*b.Location)[b.ExtractionPointer]
	b.ExtractionPointer++

	return data
}

func (b *Buffer) Write(data byte) {
	b.Location.Append(data)
	b.ExtractionPointer++
}

func (b *Buffer) Reset() {
	b.ExtractionPointer = 0
}
