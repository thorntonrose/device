package device

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thorntonrose/device/internal/config"
	"github.com/thorntonrose/device/internal/etc"
)

const LOG_FILE = "device.log"

func Run() {
	defer config.InitLogger()()

	slot := flag.Int("slot", 0, "")
	flag.Usage = Usage
	flag.Parse()

	device := New()
	device.Load(string(etc.Must(os.ReadFile(GetFile()))))
	device.Run(*slot)

	// ???: Dump non-empty only?
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
