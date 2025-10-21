package bins

import (
	"encoding/json"
	"fmt"
	promptdata "json/promptData"
	"time"
)

type Bin struct {
	ID        string    `id:"id"`
	Name      string    `id:"name"`
	Private   bool      `id:"private"`
	CreatedAt time.Time `id:"created"`
	UpdatedAt time.Time `id:"updated"`
}

type BinList struct {
	Accounts  []Bin     `id:"accounts"`
	UpdatedAt time.Time `id:"updated"`
}

type DB interface {
	Read() ([]byte, error)
	Write(content []byte)
}

type BinListWithDb struct {
	BinList
	db DB
}

func CreateBin(db DB) (binListWithDb *BinListWithDb, err error) {
	data, err := db.Read()
	if err != nil {
		db.Write(data)
		return &BinListWithDb{
			BinList: BinList{
				Accounts:  []Bin{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}, nil
	}
	var binList BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		fmt.Println("Ошибка преобразования в struct!")
	}
	bin, err := NewBin()
	if err != nil {
		fmt.Println(err)
	}
	binList.AddInBinList(bin)
	content, err := binList.ToByte()
	if err != nil {
		fmt.Println(err)
	}
	db.Write(content)
	return &BinListWithDb{
		BinList: BinList{
			Accounts:  binList.Accounts,
			UpdatedAt: time.Now(),
		},
		db: db,
	}, nil
}

func (binList *BinList) AddInBinList(bin *Bin) {
	binList.Accounts = append(binList.Accounts, *bin)
	binList.UpdatedAt = time.Now()
}

func (binList *BinList) ToByte() ([]byte, error) {
	data, err := json.Marshal(binList)
	if err != nil {
		promptdata.PrintErr("Ошибка преобразования в JSON!")
		return nil, err
	}
	return data, nil
}

func NewBin() (*Bin, error) {
	id, err := promptdata.PromptData("Введите ID")
	if err != nil {
		promptdata.PrintErr("Ошибка ввода!")
		return nil, err
	}
	name, err := promptdata.PromptData("Введите имя пользователя")
	if err != nil {
		promptdata.PrintErr("Ошибка ввода!")
		return nil, err
	}
	var privateChoise int
	var private bool
	fmt.Println("Выберете приватность профиля: 1 - приватный\n2 - нет")
	fmt.Println()
	fmt.Scan(&privateChoise)
	switch privateChoise {
	case 1:
		private = true
	case 2:
		private = false
	default:
		fmt.Println("Ошибка ввода!")
		return nil, err
	}
	return &Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
