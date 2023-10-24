package util

import (
	"runtime"
)

func GetContainerIP() string {
	os := runtime.GOOS

	if os == "windows" {
		return "postgresql://root:secret@172.29.46.53:5432/simple_bank?sslmode=disable"
	} else {
		return "postgresql://root:secret@0.0.0.0:5432/simple_bank?sslmode=disable"
	}
}
