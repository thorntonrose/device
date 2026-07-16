package device

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/thorntonrose/device/device/config"
	. "github.com/thorntonrose/device/device/etc"
)

type Flags struct {
	Dump  *bool
	Log   *string
	Model *int
}

func Run() {
	flags := ParseFlags()
	config.InitNewLog(*flags.Log)
	RunProgram(flags)
}

func ParseFlags() (flags Flags) {
	flags.Dump = flag.Bool("dump", false, "")
	flags.Log = flag.String("log", "", "")
	flags.Model = flag.Int("model", config.MaxModels, "")
	flag.Usage = Usage
	flag.Parse()

	return
}

func Usage() {
	fmt.Printf("Usage: %s [flags] <file>\n", filepath.Base(os.Args[0]))
	fmt.Println("Flags:")
	fmt.Println("  -dump = dump memory at end of program")
	fmt.Printf("  -model <number> = device model (default: %d)\n", config.MaxModels)
	fmt.Println("  -log <file> = log to <file>")

	os.Exit(1)
}

func RunProgram(flags Flags) {
	config.Model = *flags.Model
	log.Printf("Model: %d\n", config.Model)

	device := New()
	device.Load(string(Must(os.ReadFile(GetFile()))))
	device.Run()

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
