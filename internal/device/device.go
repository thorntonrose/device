package device

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thorntonrose/device/internal/buf"
	"github.com/thorntonrose/device/internal/mem"
)

var DirectivePattern = regexp.MustCompile(`^(\d)+(=|\$)(.*)$`)
var CommandPattern = regexp.MustCompile(`^(\*|\+)?([A-Z])([0-9\.])*`)

func Run() {
	device := New()
	fmt.Println(device.Memory.Dump(mem.Transmit, 3, 20))

	device.Load("003=HELLO\n020$X")
	fmt.Println(device.Memory.Dump(mem.Transmit, 3, 20))

	device.Run(20)
	fmt.Println(device.Memory.Dump(mem.Transmit, 3, 20))
}

//-----------------------------------------------------------------------------

type Device struct {
	Memory    mem.Memory
	BufferSet buf.BufferSet
	Commands  map[string]func(parameters []string)
}

func New() Device {
	device := Device{Memory: mem.New()}
	device.BufferSet = buf.NewBufferSet(device.Memory, 0)
	device.Commands = map[string]func(parameters []string){"X": device.RunX}

	return device
}

func (d Device) Load(program string) {
	for _, line := range strings.Split(program, "\n") {
		d.Memory.Set(d.Parse(line))
	}
}

func (d Device) Parse(line string) (location int, value []byte) {
	// expect: [<line>, <location>, <operator>, <value>]
	if matches := DirectivePattern.FindStringSubmatch(line); len(matches) == 4 {
		location, _ := strconv.Atoi(matches[1])
		return location, []byte(matches[3])
	}

	panic(fmt.Sprintf("invalid directive: %s", line))
}

func (d Device) Run(location int) {
	script := string(d.Memory[location])

	for _, tokens := range CommandPattern.FindAllStringSubmatch(script, -1) {
		d.RunCommand(tokens)
	}
}

func (d Device) RunCommand(tokens []string) {
	d.Commands[tokens[1]+tokens[2]](tokens[3:])
}

func (d Device) RunX(parameters []string) {
	d.BufferSet.Copy()
}
