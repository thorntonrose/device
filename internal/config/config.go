package config

import (
	"log"
	"os"

	"github.com/thorntonrose/device/internal/etc"
)

var LogFile = GetEnv("LOG_FILE", "device.log")

func GetEnv(key, def string) string {
	return etc.Value(os.Getenv(key), def)
}

func InitLogger() func() {
	writer := etc.Must(os.OpenFile(LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644))
	log.SetOutput(writer)

	return func() { writer.Close() }
}
