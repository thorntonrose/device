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
	Version = "0.1.0"

	BinDir      = "bin"
	DistFile    = Name + ".tgz"
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

	for _, path := range []string{BinDir, TempDir, DistFile} {
		os.RemoveAll(path)
	}

	Bash("go clean -cache")
}

func Build() {
	fmt.Println("> Build")

	sh.Rm(BinDir)
	build("", "")
	build("linux", "amd64")
	build("darwin", "arm64")
}

func build(goos, goarch string) {
	Env["GOOS"] = goos
	Env["GOARCH"] = goarch
	name := strings.Replace(fmt.Sprintf("%s.%s.%s", Name, goos, goarch), "..", "", 1)

	Bash("go build -ldflags '-X main.Version=%s' -o ./%s/%s ./cmd", Version, BinDir, name)
}

func Dist() {
	Bash("tar -czf %s docs/* bin/*", DistFile)
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
