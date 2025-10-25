package api

import "json/config"

func GetApiKey() string {
	cfg := config.NewConfig()
	return cfg.Key
}
