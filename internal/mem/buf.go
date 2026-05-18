package mem

import "github.com/thorntonrose/device/internal/iter"

func (self *Memory) ReadAll(bufNum int, maxCount int, stop byte) (data []byte) {
	for count := 0; self.HasNext(bufNum, count, maxCount); count++ {
		b := self.Read(bufNum)

		if stop > 0 && b == stop {
			break
		}

		data = append(data, b)
	}

	return
}

func (self *Memory) HasNext(bufNum int, count int, maxCount int) bool {
	return (maxCount == 0 || count < maxCount) && self.Pointers[bufNum] < len(self.Slots[bufNum+1])
}

func (self *Memory) Read(bufNum int) byte {
	data := self.Slots[bufNum+1][self.Pointers[bufNum]]
	self.Pointers[bufNum]++

	return data
}

func (self *Memory) WriteAll(bufNum int, data []byte) {
	iter.Each(data, func(b byte) { self.Write(bufNum, b) })
}

func (self *Memory) Write(bufNum int, data byte) {
	self.Slots[bufNum+1] = append(self.Slots[bufNum+1], data)
	self.Pointers[bufNum]++
}

func (self *Memory) Clear(bufNum int) {
	self.Slots[bufNum+1] = self.Slots[bufNum+1][:0]
	self.Reset(bufNum)
}

func (self *Memory) Reset(bufNum int) {
	self.Pointers[bufNum] = 0
}

func (self *Memory) Trim(bufNum int, n int) {
	self.Pointers[bufNum] = max(self.Pointers[bufNum]-n, 0)
	self.Slots[bufNum+1] = self.Slots[bufNum+1][:self.Pointers[bufNum]]
}
