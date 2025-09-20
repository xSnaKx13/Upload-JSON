package main

import (
	"errors"
	"fmt"
	"json/bins"
	"time"
)

func main() {
	newBin, err := constructBin("231433312", true, "dea")
	if err != nil {
		fmt.Print("Ошибка ", err)
	}
	fmt.Print(newBin)
}

func constructBin(id string, private bool, name string) (*bins.Bin, error) {
	if id == "" || len(id) == 0 {
		return nil, errors.New("ID не может быть пустым")
	}
	if name == "" || len(name) == 0 {
		return nil, errors.New("NAME не может быть пустым")
	}
	return &bins.Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}, nil
}
