package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteFile(b []byte) {
	fmt.Print("Введите название файла: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	name = strings.TrimSpace(name)

	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Ошибка создания файла!")
		return
	}
	_, err = file.Write(b)
	if err != nil {
		fmt.Println("Ошибка записи!")
		return
	}
	fmt.Println("Запись прошла успешно.")
}

func ReadFile(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
