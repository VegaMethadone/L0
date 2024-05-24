package client

/*
import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func Client(message string) {
	// Подключение к NATS серверу
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS серверу: %v", err)
	}
	defer nc.Close()

	// Подключение к NATS Streaming серверу
	sc, err := stan.Connect("test-cluster", "client-456", stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS Streaming серверу: %v", err)
	}
	defer sc.Close()

	// Публикация сообщения в канал "subject-name"
	err = sc.Publish("subject-name", []byte(message))
	if err != nil {
		log.Fatalf("Ошибка публикации сообщения: %v", err)
	}

	log.Println("Сообщение опубликовано:", message)
}
*/
import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

// Структура для примера
type Message struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func Client(message Message) {
	// Подключение к NATS серверу
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS серверу: %v", err)
	}
	defer nc.Close()

	// Подключение к NATS Streaming серверу
	sc, err := stan.Connect("test-cluster", "client-456", stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS Streaming серверу: %v", err)
	}
	defer sc.Close()

	// Сериализация структуры в JSON
	data, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Ошибка сериализации JSON: %v", err)
	}

	// Публикация сообщения в канал "subject-name"
	err = sc.Publish("subject-name", data)
	if err != nil {
		log.Fatalf("Ошибка публикации сообщения: %v", err)
	}

	log.Println("Сообщение опубликовано:", string(data))
}
