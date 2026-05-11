package mem

import "github.com/thorntonrose/device/internal/iter"

func (m *Memory) ReadAll(bufNum int, maxCount int, stop byte) (data []byte) {
	for count := 0; m.HasNext(bufNum, count, maxCount); count++ {
		b := m.Read(bufNum)

		if stop > 0 && b == stop {
			break
		}

		data = append(data, b)
	}

	return
}

func (m *Memory) HasNext(bufNum int, count int, maxCount int) bool {
	return (maxCount == 0 || count < maxCount) && m.Pointers[bufNum] < len(m.Slots[bufNum+1])
}

func (m *Memory) Read(bufNum int) byte {
	data := m.Slots[bufNum+1][m.Pointers[bufNum]]
	m.Pointers[bufNum]++

	return data
}

func (m *Memory) WriteAll(bufNum int, data []byte) {
	iter.Each(data, func(b byte) { m.Write(bufNum, b) })
}

func (m *Memory) Write(bufNum int, data byte) {
	m.Slots[bufNum+1] = append(m.Slots[bufNum+1], data)
	m.Pointers[bufNum]++
}

func (m *Memory) Clear(bufNum int) {
	m.Slots[bufNum+1] = m.Slots[bufNum+1][:0]
	m.Reset(bufNum)
}

func (m *Memory) Reset(bufNum int) {
	m.Pointers[bufNum] = 0
}
