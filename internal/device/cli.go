package device

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thorntonrose/device/internal/etc"
)

func Run() {
	slot := flag.Int("slot", 0, "")
	flag.Usage = Usage
	flag.Parse()
	fileName := GetFile()

	device := New()
	device.Load(string(etc.Must(os.ReadFile(fileName))))
	device.Run(*slot)
	fmt.Println(device.Memory.Dump())
}

func Usage() {
	fmt.Printf("Usage: %s [flags] <file>\n", filepath.Base(os.Args[0]))
	fmt.Println("Flags:")
	fmt.Println("  -slot = script slot number")

	os.Exit(1)
}

func GetFile() string {
	if flag.NArg() < 1 {
		Usage()
	}

	return flag.Arg(0)
}
