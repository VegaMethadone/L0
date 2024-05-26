package client

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func Client(message Order) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS серверу: %v", err)
	}
	defer nc.Close()

	sc, err := stan.Connect("test-cluster", "client-456", stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS Streaming серверу: %v", err)
	}
	defer sc.Close()

	data, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Ошибка сериализации JSON: %v", err)
	}

	err = sc.Publish("subject-name", data)
	if err != nil {
		log.Fatalf("Ошибка публикации сообщения: %v", err)
	}

	log.Println("Сообщение опубликовано:", string(data))
}
