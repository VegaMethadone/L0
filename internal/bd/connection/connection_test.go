package connection

import (
	"fmt"
	"log"
	"testing"
)

func TestGetConnection(t *testing.T) {
	result := getConnectionStr()
	if result == "" {
		log.Fatalf("Condif is damaged")
		return
	}
	fmt.Println(result)
}
