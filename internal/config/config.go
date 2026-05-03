package config

import (
	"io"
	"log"
	"os"

	"github.com/thorntonrose/device/internal/etc"
)

var LogFile = GetEnv("LOG_FILE", "device.log")

//-----------------------------------------------------------------------------

func GetEnv(key, def string) string {
	return etc.Value(os.Getenv(key), def)
}

//-----------------------------------------------------------------------------

func InitNewLog() func() {
	os.Remove(LogFile)
	return InitLog()
}

func InitLog() func() {
	return etc.If(LogFile == "none", InitNopLog, InitFileLog)()
}

func InitNopLog() func() {
	log.SetOutput(io.Discard)
	return func() {}
}

func InitFileLog() func() {
	writer := etc.Must(os.OpenFile(LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644))
	log.SetOutput(writer)

	return func() { writer.Close() }
}
