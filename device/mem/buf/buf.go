package buf

import (
	"github.com/thorntonrose/device/device/iter"
	"github.com/thorntonrose/device/device/mem"
)

func ReadAll(memory *mem.Memory, bufNum int, maxCount int, stop byte) (data []byte) {
	for count := 0; HasNext(memory, bufNum, count, maxCount); count++ {
		b := Read(memory, bufNum)

		if stop > 0 && b == stop {
			break
		}

		data = append(data, b)
	}

	return
}

func HasNext(memory *mem.Memory, bufNum int, count int, maxCount int) bool {
	return (maxCount == 0 || count < maxCount) && memory.Pointers[bufNum] < len(memory.Slots[bufNum+1])
}

func Read(memory *mem.Memory, bufNum int) (data byte) {
	data = memory.Slots[bufNum+1][memory.Pointers[bufNum]]
	memory.Pointers[bufNum]++

	return
}

//-----------------------------------------------------------------------------

func WriteAll(memory *mem.Memory, bufNum int, data []byte) {
	iter.Each(data, func(b byte) { Write(memory, bufNum, b) })
}

func Write(memory *mem.Memory, bufNum int, data byte) {
	memory.Slots[bufNum+1] = append(memory.Slots[bufNum+1], data)
	memory.Pointers[bufNum]++
}

//-----------------------------------------------------------------------------

func Clear(memory *mem.Memory, bufNum int) {
	memory.Slots[bufNum+1] = memory.Slots[bufNum+1][:0]
	Reset(memory, bufNum)
}

func Reset(memory *mem.Memory, bufNum int) {
	memory.Pointers[bufNum] = 0
}
