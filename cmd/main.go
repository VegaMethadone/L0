package main

import (
	"L0/internal/bd"
	"L0/internal/server"
	"L0/internal/server/cache"
	"L0/internal/structs"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func init() {
	cache.Cache = cache.NewRwCache()
	err := cache.Cache.Restore()
	if err != nil {
		log.Fatalf("Could not restore  cache: %v", err)
		return
	}
	log.Println("cache is restored")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		connectAndListenNats()
	}()

	go func() {
		defer wg.Done()
		startServer()
	}()

	wg.Wait()
}

func connectAndListenNats() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error connecting to NATS server: %v", err)
	}
	defer nc.Close()

	sc, err := stan.Connect("test-cluster", "client-123", stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Error connecting to NATS Streaming server: %v", err)
	}
	defer sc.Close()

	natsMessageHandler := func(msg *stan.Msg) {
		fmt.Println("Received message")

		var gotData structs.Order
		if err := json.Unmarshal(msg.Data, &gotData); err != nil {
			log.Printf("Error decoding message: %v\n", err)
			return
		}

		if cache.Cache.IsExist(gotData.OrderUID) {
			return
		}

		if err := bd.PostData(string(msg.Data)); err != nil {
			log.Printf("Error posting data: %v\n", err)
			return
		}

		cache.Cache.Add(gotData.OrderUID, &gotData)
	}

	sub, err := sc.Subscribe("subject-name", natsMessageHandler, stan.DeliverAllAvailable())
	if err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("Subscribed to the channel. Waiting for messages...")
	select {}
}

func startServer() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("Faild to start server: %v\n", err)
		return
	}

	log.Printf("Server is working at: http://%s\n", srv.Addr)
	log.Printf("if you want to check order: http://%s/orders/{id}", srv.Addr)
	log.Panicln(srv.ListenAndServe())
}
