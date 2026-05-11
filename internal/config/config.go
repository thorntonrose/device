package config

import (
	"io"
	"log"
	"os"

	. "github.com/thorntonrose/device/internal/etc"
)

var LogFile = GetEnv("LOG_FILE", "")
var LogWriter *os.File

//-----------------------------------------------------------------------------

func GetEnv(key, def string) string {
	return Value(os.Getenv(key), def)
}

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
	LogWriter = Must(os.OpenFile(LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644))
	log.SetOutput(LogWriter)

	return func() { LogWriter.Close() }
}
