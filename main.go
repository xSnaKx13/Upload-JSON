package main

import (
	"fmt"
	"json/bins"
	"json/file"
)

func main() {
	db := file.NewJsonDb("data.json")
	_, err := bins.CreateBin(db)
	if err != nil {
		fmt.Println(err)
	}

}
