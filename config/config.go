package config

import (
	"encoding/json"
	"os"
)

// Определение структуры для конфигурации
type Config struct {
	Version int    `json:"version"`
	Env     string `json:"env"`
	Network struct {
		Address      string `json:"address"`
		Port         string `json:"port"`
		WriteTimeout int    `json:"writeTimeout"`
		ReadTimeout  int    `json:"readTimeout"`
	} `json:"network"`
	Postgres struct {
		Host         string `json:"host"`
		Port         string `json:"port"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		DatabaseName string `json:"database_name"`
		Sslmode      string `json:"sslmode"`
	} `json:"postgres"`
}

// Функция для получения конфигурации из файла
func GetConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
