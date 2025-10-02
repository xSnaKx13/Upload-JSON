package main

import (
	"fmt"
	"json/bins"
	"json/storage"
)

func main() {
	newBin, err := bins.CreateNewBin()
	if err != nil {
		fmt.Println(err)
		return
	}
	binList := bins.CnstructBinList()
	binList.AddBins(newBin)
	file, err := binList.ToByte(binList)
	if err != nil {
		fmt.Println(err)
		return
	}
	storage.WriteFileInJson(file)
}
