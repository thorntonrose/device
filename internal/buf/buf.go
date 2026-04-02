package buf

import (
	"github.com/thorntonrose/device/internal/mem"
)

const (
	MinBuffers = 2

	Transmit = 0
	Receive  = 1
)

type BufferSet struct {
	Buffers     []Buffer
	Source      *Buffer
	Destination *Buffer
}

func NewBufferSet(memory mem.Memory, extra int) BufferSet {
	buffers := make([]Buffer, 0, MinBuffers+extra)
	buffers = append(buffers, NewBuffer(memory, mem.Transmit))
	buffers = append(buffers, NewBuffer(memory, mem.Receive))

	for i := 0; i < extra; i++ {
		buffers = append(buffers, NewBuffer(memory, MinBuffers+i))
	}

	return BufferSet{Buffers: buffers, Source: &buffers[Receive], Destination: &buffers[Transmit]}
}

func (b BufferSet) Copy() {
	for i := b.Source.Pointer; i < len(b.Source.Memory[b.Source.Location]); i++ {
		b.Destination.Write(b.Source.Read())
	}
}

func (b BufferSet) Read() byte {
	return b.Source.Read()
}

func (b BufferSet) Write(data byte) {
	b.Destination.Write(data)
}

//-----------------------------------------------------------------------------

type Buffer struct {
	Memory   mem.Memory
	Location int
	Pointer  int
}

func NewBuffer(memory mem.Memory, location int) Buffer {
	return Buffer{Memory: memory, Location: location}
}

func (b Buffer) Get() []byte {
	return b.Memory[b.Location]
}

func (b *Buffer) Set(data []byte) {
	b.Memory[b.Location] = data
	b.Pointer = len(data)
}

func (b *Buffer) Read() byte {
	data := b.Memory[b.Location][b.Pointer]
	b.Pointer++

	return data
}

func (b *Buffer) Write(data byte) {
	b.Memory.Append(b.Location, data)
	b.Pointer++
}

func (b *Buffer) Reset() {
	b.Pointer = 0
}
