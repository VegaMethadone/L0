package bd

import (
	"L0/internal/bd/query"
	"L0/internal/structs"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestPostData(t *testing.T) {
	var testData = &structs.Order{
		OrderUID:    "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: structs.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: structs.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDT:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []structs.Item{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				RID:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NMID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
			{
				ChrtID:      9934931,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				RID:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NMID:        2389212,
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
		DateCreated:       time.Date(2021, time.November, 26, 6, 22, 19, 0, time.UTC),
		OofShard:          "1",
	}

	jsonData, err := json.Marshal(testData)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	fmt.Println(string(jsonData))
	err = PostData(string(jsonData))
	if err != nil {
		log.Fatalf("Error add datato BD: %v", err)
	}
}

func TestGetData(t *testing.T) {
	id := "b563feb7b2b84b6test"
	data, err := GetData(id)
	if err != nil {
		log.Fatalf("Could not get data: %v", data)
	}
	prettyPrint(data)
}

func TestCleatAllData(t *testing.T) {
	id := "b563feb7b2b84b6test"
	query.CleanData(id)
}

func prettyPrint(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("Could not format JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}
