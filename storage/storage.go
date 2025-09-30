package storage

import (
	"json/file"
)

func WriteFileInJson(b []byte) {
	file.WriteFile(b)
}
func ReadJsonFile(name string) {
	file.ReadFile(name)
}
