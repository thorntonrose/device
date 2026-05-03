package io

import (
	"testing"

	"github.com/thorntonrose/device/internal/config"
)

func TestMain(m *testing.M) {
	defer config.InitLog()()
	m.Run()
}
