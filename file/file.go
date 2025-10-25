package file

import (
	promptdata "json/promptData"
	"os"
	"strings"
)

func Read(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		promptdata.PrintErr("Ошибка чтения файла!")
		return nil, err
	}
	if !strings.Contains(string(data), "json") {
		promptdata.PrintErr("Файл не содержит 'json'!")
		return nil, err
	}
	return data, nil
}
