package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	newBin, err := constructBin("231433312", true, "")
	if err != nil {
		fmt.Print("Ошибка ", err)
	}
	fmt.Print(newBin)
}

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	Bins []*Bin
}

func constructBin(ID string, Private bool, Name string) (*Bin, error) {
	if len(ID) == 0 {
		return nil, errors.New("ID не может быть пустым")
	}
	if len(Name) == 0 {
		return nil, errors.New("Name не может быть пустым")
	}
	return &Bin{
		id:        ID,
		private:   Private,
		createdAt: time.Now(),
		name:      Name,
	}, nil
}
