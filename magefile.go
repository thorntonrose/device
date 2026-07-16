//go:build mage

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Name = "device"
	Pkg  = "./..."

	BinDir      = "bin"
	CmdDir      = "cmd"
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
	os.RemoveAll(BinDir)
	os.RemoveAll(TempDir)
	Bash("go clean -cache")
}

func Tidy() {
	fmt.Println("> Tidy")
	Bash("go mod tidy -v")
}

func Build() {
	fmt.Println("> Build")
	os.MkdirAll(BinDir, 0755)
	Bash("go build -o %s/%s ./%s", BinDir, Name, CmdDir)
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

func Run(args string) {
	mg.SerialDeps(Tidy)

	fmt.Println("> Run")
	Bash("go run ./%s/main.go %s", CmdDir, args)
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
