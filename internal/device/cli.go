package device

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thorntonrose/device/internal/config"
	"github.com/thorntonrose/device/internal/etc"
)

type Flags struct {
	Dump *bool
	Log  *string
	Slot *int
}

func Run() {
	flags := ParseFlags()
	config.LogFile = *flags.Log
	defer config.InitNewLog()()

	RunProgram(flags)
}

func ParseFlags() (flags Flags) {
	flags.Dump = flag.Bool("dump", false, "dump memory")
	flags.Log = flag.String("log", "none", "log file")
	flags.Slot = flag.Int("slot", 0, "")
	flag.Usage = Usage
	flag.Parse()

	return
}

func Usage() {
	fmt.Printf("Usage: %s [flags] <file>\n", filepath.Base(os.Args[0]))
	fmt.Println("Flags:")
	fmt.Println("  -dump = dump memory at end of program")
	fmt.Println("  -log <file> = log file")
	fmt.Println("  -slot <number> = script slot number (default: 0)")

	os.Exit(1)
}

func RunProgram(flags Flags) {
	device := New()
	device.Load(string(etc.Must(os.ReadFile(GetFile()))))
	device.Run(*flags.Slot)

	if *flags.Dump {
		fmt.Println("\n-----\n" + device.Memory.Dump())
	}
}

func GetFile() string {
	if flag.NArg() < 1 {
		Usage()
	}

	return flag.Arg(0)
}
