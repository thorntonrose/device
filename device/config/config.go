package config

import (
	"io"
	"log"
	"os"

	. "github.com/thorntonrose/device/device/etc"
)

const MaxModels = 5

var (
	LogFile = GetEnv("LOG_FILE", "")
	Model   = MaxModels
)

//-----------------------------------------------------------------------------

func InitNewLog(fileName string) func() {
	if LogFile = fileName; LogFile != "" {
		os.Remove(LogFile)
	}

	return InitLog()
}

func InitLog() func() {
	return If(LogFile == "", InitLogNop, InitLogFile)()
}

func InitLogNop() func() {
	log.SetOutput(io.Discard)
	return func() {}
}

func InitLogFile() func() {
	file := Must(os.OpenFile(LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644))
	log.SetOutput(file)

	return func() { file.Close() }
}

//-----------------------------------------------------------------------------

func GetEnv(key, def string) string {
	return Value(os.Getenv(key), def)
}
