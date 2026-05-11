package device

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thorntonrose/device/internal/config"
	. "github.com/thorntonrose/device/internal/etc"
)

var Version = "0.0.0"

type Flags struct {
	Dump    *bool
	Log     *string
	Slot    *int
	Version *bool
}

func Run() {
	flags := ParseFlags()

	if *flags.Version {
		fmt.Println("device", Version)
		os.Exit(0)
	}

	config.InitNewLog(*flags.Log)
	RunProgram(flags)
}

func ParseFlags() (flags Flags) {
	flags.Dump = flag.Bool("dump", false, "")
	flags.Log = flag.String("log", "", "")
	flags.Slot = flag.Int("slot", 0, "")
	flags.Version = flag.Bool("version", false, "")
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
	fmt.Println("  -version = print version then exit")

	os.Exit(1)
}

func RunProgram(flags Flags) {
	device := New()
	device.Load(string(Must(os.ReadFile(GetFile()))))
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
