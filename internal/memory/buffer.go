package memory

const (
	MaxBuffers    = 2
	MaxBufferSize = 250

	TransmitBufferNum = 1
	ReceiveBufferNum  = 2
)

type Buffer struct {
	Location          *[]byte
	ExtractionPointer int
}

func NewBuffers() []Buffer {
	buffers := make([]Buffer, MaxBuffers)

	for i := range buffers {
		buffers[i] = Buffer{}
	}

	return buffers
}
