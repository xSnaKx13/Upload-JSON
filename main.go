package main

import (
	"fmt"
	"json/bins"
	"json/storage"
)

func main() {
	db := storage.NewJsonDb("data.json")
	_, err := bins.CreateBin(db)
	if err != nil {
		fmt.Println(err)
	}

}
