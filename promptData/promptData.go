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
	switch v := value.(type) {
	case string:
		fmt.Println("Ошибка:", v)
	case int:
		fmt.Printf("Код ошибки %d", v)
	case error:
		fmt.Println("Ошибка:", v.Error())
	default:
		fmt.Println("Неизвестная ошибка!")
	}
}
