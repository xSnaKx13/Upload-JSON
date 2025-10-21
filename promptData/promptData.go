package promptdata

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptData(prompt string) (string, error) {
	fmt.Println(prompt + ": ")
	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	data = strings.TrimSpace(data)
	if err != nil {
		fmt.Println("Ошибка ввода!")
		return "", err
	}
	return data, nil
}

func PrintErr(value any) {
	switch t := value.(type) {
	case string:
		fmt.Println(t)
	case int:
		fmt.Printf("Код ошибки %d", t)
	case error:
		fmt.Println(t.Error())
	default:
		fmt.Println("Неизвестная ошибка!")
	}
}
