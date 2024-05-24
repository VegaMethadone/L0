package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {

	// Подключение к NATS серверу
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS серверу: %v", err)
	}
	defer nc.Close()

	// Подключение к NATS Streaming серверу
	sc, err := stan.Connect("test-cluster", "client-123", stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS Streaming серверу: %v", err)
	}
	defer sc.Close()

	// Подписка на канал
	sub, err := sc.Subscribe("subject-name", func(msg *stan.Msg) {
		fmt.Printf("Получено сообщение: %s\n", string(msg.Data))
	}, stan.DeliverAllAvailable()) // Начать считывание с самого начала
	if err != nil {
		log.Fatalf("Ошибка подписки на канал: %v", err)
	}
	defer sub.Unsubscribe()

	// Ожидание завершения
	fmt.Println("Подписан на канал. Ожидание сообщений...")
	select {} // Блокировка основной горутины, чтобы программа не завершилась
}
