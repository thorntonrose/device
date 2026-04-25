package device

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thorntonrose/device/internal/config"
	"github.com/thorntonrose/device/internal/etc"
)

func Run() {
	defer config.InitLogger()()

	slot := flag.Int("slot", 0, "")
	dump := flag.Bool("dump", false, "dump memory after running")
	flag.Usage = Usage
	flag.Parse()

	RunProgram(*slot, *dump)
}

func Usage() {
	fmt.Printf("Usage: %s [flags] <file>\n", filepath.Base(os.Args[0]))
	fmt.Println("Flags:")
	fmt.Println("  -slot = script slot number (default: 0)")
	fmt.Println("  -dump = dump memory after running")

	os.Exit(1)
}

func GetFile() string {
	if flag.NArg() < 1 {
		Usage()
	}

	return flag.Arg(0)
}

func RunProgram(slot int, dump bool) {
	device := New()
	device.Load(string(etc.Must(os.ReadFile(GetFile()))))
	device.Run(slot)

	if dump {
		fmt.Println(device.Memory.Dump())
	}
}
