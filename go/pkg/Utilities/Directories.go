package Utilities

import (
	"path"
	"runtime"
)

func GetDayDirectory() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
