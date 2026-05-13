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
	Pkg         = "./..."
	TempDir     = "tmp"
	Tests       = "Test"
	TestTempDir = TempDir + "/test"
	WorkDir, _  = os.Getwd()

	Env = map[string]string{}
)

//-----------------------------------------------------------------------------

func Clean() {
	fmt.Println("> Clean")
	os.RemoveAll(Name)
	os.RemoveAll(TempDir)
	Bash("go clean -cache")
}

func Build() {
	fmt.Println("> Build")
	Bash("go build -o %s .", Name)
}

func Test() {
	fmt.Println("> Test")
	logDir := WorkDir + "/" + TestTempDir
	Env["LOG_FILE"] = logDir + "/test.log"
	Env["TEST_TEMP_DIR"] = WorkDir + "/" + TestTempDir

	sh.Rm(TestTempDir)
	os.MkdirAll(TestTempDir, 0755)
	Bash("go test -v -tags test -run %s %s", GetEnv("TESTS", Tests), GetEnv("PACKAGE", Pkg))
}

//-----------------------------------------------------------------------------

func Bash(format string, args ...any) {
	cmd := []string{"-o", "pipefail", "-c", strings.Trim(fmt.Sprintf(format, args...), " ")}
	fmt.Println(cmd[len(cmd)-1])

	if err := sh.RunWithV(Env, "bash", cmd...); err != nil {
		os.Exit(1)
	}
}

func GetEnv(key string, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}
