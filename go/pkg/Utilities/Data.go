package Utilities

import (
	"os"
	"path"
)

func GetAllRealData(currentDirectory string) string {
	return GetAllDataAsString(path.Join(currentDirectory, "data.txt"))
}

func GetAllTestData(currentDirectory string) string {
	return GetAllDataAsString(path.Join(currentDirectory, "test_data.txt"))
}

func GetAllDataAsString(path string) string {
	dat, err := os.ReadFile(path)
	CheckError(err)
	return string(dat)
}
