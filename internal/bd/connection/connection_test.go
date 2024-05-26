package connection

import (
	"fmt"
	"log"
	"testing"
)

func TestGetConnection(t *testing.T) {
	result := getConnectionStr()
	if result == "" {
		log.Fatalf("Config is damaged")
		return
	}
	fmt.Println(result)
}
