package main

import "github.com/thorntonrose/device/internal/device"

var Version = "0.0.0"

func main() {
	device.Version = Version
	device.Run()
}
