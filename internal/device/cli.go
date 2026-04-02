package device

import (
	"fmt"

	"github.com/thorntonrose/device/internal/mem"
)

func Run() {
	device := New()
	fmt.Println(device.Memory.Dump(mem.Transmit, 3, 20))

	device.Load("003=HELLO\n020$X")
	fmt.Println(device.Memory.Dump(mem.Transmit, 3, 20))

	device.Run(20)
	fmt.Println(device.Memory.Dump(mem.Transmit, 3, 20))
}
