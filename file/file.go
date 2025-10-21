package file

import (
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

func (db JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.Name)
	if err != nil {
		promptdata.PrintErr("Ошибка чтения файла!")
		return nil, err
	}
	return data, nil
}

func (db JsonDb) Write(content []byte) {
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
