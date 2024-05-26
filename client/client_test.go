package client

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func randomString(length int) string {
	charset := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func getRandomData() *Order {
	return &Order{
		OrderUID:    randomString(20),
		TrackNumber: randomString(12),
		Entry:       "WBIL",
		Delivery: Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: Payment{
			Transaction:  randomString(20),
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       rand.Intn(5000),
			PaymentDT:    time.Now().Unix(),
			Bank:         "alpha",
			DeliveryCost: rand.Intn(2000),
			GoodsTotal:   rand.Intn(1000),
			CustomFee:    0,
		},
		Items: []Item{
			{
				ChrtID:      rand.Intn(100000),
				TrackNumber: randomString(12),
				Price:       rand.Intn(1000),
				RID:         randomString(20),
				Name:        "Mascaras",
				Sale:        rand.Intn(50),
				Size:        "0",
				TotalPrice:  rand.Intn(500),
				NMID:        rand.Intn(100000),
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
			{
				ChrtID:      rand.Intn(100000),
				TrackNumber: randomString(12),
				Price:       rand.Intn(1000),
				RID:         randomString(20),
				Name:        "Mascaras",
				Sale:        rand.Intn(50),
				Size:        "0",
				TotalPrice:  rand.Intn(500),
				NMID:        rand.Intn(100000),
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SMID:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
	}
}

func TestClient(t *testing.T) {
	for i := 0; i < 100; i++ {
		validData := getRandomData()

		data, err := json.MarshalIndent(validData, "", "	")
		if err != nil {
			log.Fatalf("Faild to convert into JSON formt: %v\n", err)
		}
		fmt.Println(string(data))
		Client(*validData)
	}
}
