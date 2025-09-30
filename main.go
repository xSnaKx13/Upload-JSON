package main

import (
	"fmt"
	"json/bins"
	"json/storage"
)

func main() {
	file, err := bins.CreateNewBin()
	if err != nil {
		fmt.Println(err)
	}
	storage.WriteFileInJson(file)
	storage.ReadJsonFile("content.json")
}
