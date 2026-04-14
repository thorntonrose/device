package config

import (
	"log"
	"os"

	"github.com/thorntonrose/device/internal/etc"
)

func GetEnv(key, def string) string {
	return etc.Value(os.Getenv(key), def)
}

func InitLogger() func() {
	file := GetEnv("LOG_FILE", "device.log")
	os.Remove(file)
	writer := etc.Must(os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644))
	log.SetOutput(writer)

	return func() { writer.Close() }
}
