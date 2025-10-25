package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("Key")
	if key == "" {
		panic("Ошибка получения ключа из переменного окружения!")
	}
	return &Config{
		Key: key,
	}
}
