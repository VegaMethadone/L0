package query

import (
	"L0/internal/bd/connection"
	"L0/internal/structs"
	"log"
)

func InsertOrderData(data *structs.Order) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	query := `
        INSERT INTO orders (
            order_uid, track_number, entry, locale, internal_signature, customer_id, 
            delivery_service, shardkey, sm_id, date_created, oof_shard
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `

	_, err = db.Exec(query,
		data.OrderUID, data.TrackNumber, data.Entry, data.Locale,
		data.InternalSignature, data.CustomerID, data.DeliveryService,
		data.ShardKey, data.SMID, data.DateCreated, data.OofShard,
	)
	if err != nil {
		log.Printf("Failed to insert order data: %v", err)
		return err
	}

	return nil
}

func InsertDeliveryData(data *structs.Order) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	query := `
        INSERT INTO delivery (
            order_uid, name, phone, zip, city, address, region, email
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	_, err = db.Exec(query,
		data.OrderUID, data.Delivery.Name, data.Delivery.Phone, data.Delivery.Zip,
		data.Delivery.City, data.Delivery.Address, data.Delivery.Region,
		data.Delivery.Email,
	)
	if err != nil {
		log.Printf("Failed to insert delivery data: %v", err)
		return err
	}

	return nil
}

func InsertPaymentData(data *structs.Order) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	query := `
        INSERT INTO payment (
            transaction, order_uid, request_id, currency, provider, amount, payment_dt,
            bank, delivery_cost, goods_total, custom_fee
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `

	_, err = db.Exec(query,
		data.Payment.Transaction, data.OrderUID, data.Payment.RequestID,
		data.Payment.Currency, data.Payment.Provider, data.Payment.Amount,
		data.Payment.PaymentDT, data.Payment.Bank, data.Payment.DeliveryCost,
		data.Payment.GoodsTotal, data.Payment.CustomFee,
	)
	if err != nil {
		log.Printf("Failed to insert payment data: %v", err)
		return err
	}

	return nil
}

func InsertItemData(id string, data *structs.Item) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	query := `
        INSERT INTO items (
            chrt_id, order_uid, track_number, price, rid, name, sale, size, total_price,
            nm_id, brand, status
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `

	_, err = db.Exec(query,
		data.ChrtID, id, data.TrackNumber, data.Price, data.RID,
		data.Name, data.Sale, data.Size, data.TotalPrice, data.NMID,
		data.Brand, data.Status,
	)
	if err != nil {
		log.Printf("Failed to insert item data: %v", err)
		return err
	}

	return nil
}
