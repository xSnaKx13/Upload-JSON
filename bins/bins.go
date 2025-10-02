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

type BinList struct {
	Bins []Bin
}

func CreateNewBin() (bin *Bin, err error) {
	bin = constructBin("33", true, "Тест")

	if bin.Id == "" {
		fmt.Println("Не корректный ID!")
		return
	}

	if bin.Name == "" {
		fmt.Println("Не корректный ID!")
		return
	}
	return bin, nil
}

func (binList *BinList) ToByte(bin *BinList) ([]byte, error) {
	file, err := json.Marshal(binList)
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

func CnstructBinList() *BinList {
	return &BinList{
		Bins: []Bin{},
	}
}

func (binList *BinList) AddBins(bin *Bin) {
	binList.Bins = append(binList.Bins, *bin)
}
