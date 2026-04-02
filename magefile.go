//go:build mage

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/sh"
)

var (
	Name    = "device"
	BinDir  = "bin"
	TempDir = "tmp"

	Pkg   = "./..."
	Tests = "Test"

	Env = map[string]string{}
)

func Clean() {
	os.RemoveAll(BinDir)
	os.RemoveAll(TempDir)
	bash("go clean -cache")
}

func Build() {
	bash("go build -o %s/%s ./cmd", BinDir, Name)
}

func Test() {
	bash("go test -v -run %s %s", getEnv("TESTS", Tests), getEnv("PACKAGE", Pkg))
}

func bash(format string, args ...any) {
	cmd := []string{"-o", "pipefail", "-c", strings.Trim(fmt.Sprintf(format, args...), " ")}
	fmt.Println(cmd[len(cmd)-1])

	if err := sh.RunWithV(Env, "bash", cmd...); err != nil {
		os.Exit(1)
	}
}

func getEnv(key string, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}
