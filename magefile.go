//go:build mage

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/sh"
)

var (
	Name        = "device"
	BinDir      = "bin"
	TempDir     = "tmp"
	TestTempDir = TempDir + "/test"
	WorkDir, _  = os.Getwd()

	Pkg   = "./..."
	Tests = "Test"

	Env = map[string]string{}
)

func Clean() {
	fmt.Println("> Clean")

	for _, path := range []string{BinDir, TempDir, "device.log"} {
		os.RemoveAll(path)
	}

	bash("go clean -cache")
}

func Build() {
	fmt.Println("> Build")
	bash("go build -o %s/%s ./cmd", BinDir, Name)
}

func Test() {
	fmt.Println("> Test")
	logDir := WorkDir + "/" + TestTempDir
	Env["LOG_FILE"] = logDir + "/test.log"
	Env["TEST_TEMP_DIR"] = WorkDir + "/" + TestTempDir

	sh.Rm(TestTempDir)
	os.MkdirAll(TestTempDir, 0755)
	bash("go test -v -tags test -run %s %s", getEnv("TESTS", Tests), getEnv("PACKAGE", Pkg))
}

//-----------------------------------------------------------------------------

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
