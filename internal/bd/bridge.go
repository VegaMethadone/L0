package bd

import (
	"L0/internal/bd/query"
	"L0/internal/structs"
	"encoding/json"
	"log"
)

func GetData(id string) (*structs.Order, error) {
	var gotData *structs.Order

	return gotData, nil
}

func PostData(data string) error {

	var orderData *structs.Order

	err := json.Unmarshal([]byte(data), &orderData)
	if err != nil {
		log.Printf("Error during unmarshal data  to json: %v", err)
		return err
	}

	err = query.InsertOrderData(orderData)
	if err != nil {
		// delete
		return err
	}
	err = query.InsertDeliveryData(orderData)
	if err != nil {
		// delete
		return err
	}
	err = query.InsertPaymentData(orderData)
	if err != nil {
		// delete
		return err
	}

	for _, value := range orderData.Items {
		err = query.InsertItemData(orderData.OrderUID, &value)
		if err != nil {
			// delete
			return err
		}
	}

	return nil
}
