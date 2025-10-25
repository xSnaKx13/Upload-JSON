package storage

import (
	"encoding/json"
	"fmt"
	promptdata "json/promptData"
	"os"
)

type JsonDb struct {
	Name string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		Name: name,
	}
}

func (db JsonDb) ReadingFromFile() ([]byte, error) {
	data, err := os.ReadFile(db.Name)
	if err != nil {
		promptdata.PrintErr("Ошибка чтения файла!")
		return nil, err
	}
	return data, nil
}

func (db JsonDb) WriteToFile(content []byte) {
	data, err := os.Create(db.Name)
	if err != nil {
		promptdata.PrintErr("Ошибка создания файла!")
		return
	}
	_, err = data.Write(content)
	if err != nil {
		promptdata.PrintErr("Ошибка записи в файл!")
		return
	}
	fmt.Println("Запись данных успешна.")
	defer data.Close()
}

func ToByte(value any) ([]byte, error) {
	data, err := json.Marshal(value)
	if err != nil {
		promptdata.PrintErr("Ошибка преобразования в JSON!")
		return nil, err
	}
	return data, nil
}
