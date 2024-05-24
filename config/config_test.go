package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	res, err := GetConfig()
	if err != nil {
		t.Errorf("Ошибка при получении конфигурации: %v", err)
		return
	}

	// Распечатываем полученную конфигурацию
	fmt.Printf("Конфигурация: %+v\n", res)
}
