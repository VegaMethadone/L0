package bd

import (
	"L0/internal/bd/query"
	"L0/internal/structs"
	"encoding/json"
	"log"
)

func GetUIDs() ([]*structs.Order, error) {
	allIds, err := query.GetAllOrderUIDs()
	if err != nil {
		log.Printf("Faild to get all UIDs: %v", err)
	}
	var arr []*structs.Order

	for _, id := range allIds {
		order, err := GetData(id)
		if err != nil {
			continue
		}
		arr = append(arr, order)
	}
	return arr, nil
}

func GetData(id string) (*structs.Order, error) {
	var gotData structs.Order

	err := query.GetOrder(id, &gotData)
	if err != nil {
		log.Printf("Error during getting  order data: %v", err)
		return nil, err
	}

	err = query.GetDelivery(id, &gotData)
	if err != nil {
		log.Printf("Error during getting delivery data: %v", err)
		return nil, err
	}

	err = query.GetPayment(id, &gotData)
	if err != nil {
		log.Printf("Error during getting payment data: %v", err)
		return nil, err
	}

	err = query.GetItems(id, &gotData)
	if err != nil {
		log.Printf("Error during getting items data: %v", err)
		return nil, err
	}

	return &gotData, nil
}

func PostData(data string) error {

	var orderData *structs.Order

	err := json.Unmarshal([]byte(data), &orderData)
	if err != nil {
		log.Printf("Error during unmarshal data  to json: %v", err)
		query.CleanData(orderData.OrderUID)
		return err
	}

	err = query.InsertOrderData(orderData)
	if err != nil {
		query.CleanData(orderData.OrderUID)
		return err
	}
	err = query.InsertDeliveryData(orderData)
	if err != nil {
		query.CleanData(orderData.OrderUID)
		return err
	}
	err = query.InsertPaymentData(orderData)
	if err != nil {
		query.CleanData(orderData.OrderUID)
		return err
	}

	for _, value := range orderData.Items {
		err = query.InsertItemData(orderData.OrderUID, &value)
		if err != nil {
			query.CleanData(orderData.OrderUID)
			return err
		}
	}

	return nil
}
