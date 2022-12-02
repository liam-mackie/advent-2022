package advent

import (
	"fmt"
	"github.com/liam-mackie/advent2022/pkg/Utilities"
	"os"
	"path"
)

func IsValidDay(day string) error {
	dir, _ := os.Getwd()
	dirList, _ := os.ReadDir(path.Join(dir, "pkg", "days"))

	validDays := make(Utilities.StringSet, 0)
	for _, entry := range dirList {
		if entry.IsDir() {
			validDays[entry.Name()] = true
		}
	}

	if validDays.Has(day) {
		return nil
	}

	err := fmt.Errorf("invalid day specified: %s \n valid days are: %s", day, validDays.Keys())
	return err
}

type Day interface {
	Run()
}
