package bins

import (
	"encoding/json"
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created"`
	Name      string    `json:"name"`
}

func CreateNewBin() ([]byte, error) {
	bin := constructBin("33", true, "Тест")

	file, err := bin.ToByte()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return file, nil
}

func (bin *Bin) ToByte() ([]byte, error) {
	file, err := json.Marshal(bin)
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return nil, err
	}
	return file, nil
}

func constructBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

type BinList struct {
	Bins []*Bin
}
