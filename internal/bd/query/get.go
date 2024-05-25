package query

import (
	"L0/internal/bd/connection"
	"L0/internal/structs"
	"database/sql"
	"errors"
	"log"
)

func GetAllOrderUIDs() ([]string, error) {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return nil, err
	}
	defer db.Close()

	query := `SELECT order_uid FROM orders`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to get all orders_id : %v", err)
		return nil, err
	}

	var orderUIDs []string
	for rows.Next() {
		var orderID string
		if err := rows.Scan(&orderID); err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		orderUIDs = append(orderUIDs, orderID)

		if err := rows.Err(); err != nil {
			log.Printf("Error occurred while iterating rows: %v", err)
			return nil, err
		}

	}
	return orderUIDs, nil
}

func GetOrder(id string, data *structs.Order) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	row := db.QueryRow(`
		SELECT
			*
		FROM
			orders
		WHERE
		order_uid = $1
	`, id)

	err = row.Scan(
		&data.OrderUID,
		&data.TrackNumber,
		&data.Entry,
		&data.Locale,
		&data.InternalSignature,
		&data.CustomerID,
		&data.DeliveryService,
		&data.ShardKey,
		&data.SMID,
		&data.DateCreated,
		&data.OofShard,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No order found with id: %s", id)
			return nil
		}
		log.Printf("Failed to scan order: %v", err)
		return err
	}

	return nil
}

func GetDelivery(id string, data *structs.Order) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	row := db.QueryRow(`
		SELECT
			*
		FROM 
			delivery
		WHERE
			order_uid = $1
	`, id)

	err = row.Scan(
		&data.OrderUID,
		&data.Delivery.Name,
		&data.Delivery.Phone,
		&data.Delivery.Zip,
		&data.Delivery.City,
		&data.Delivery.Address,
		&data.Delivery.Region,
		&data.Delivery.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No delivery found for order with id: %s", id)
			return nil
		}
		log.Printf("Failed to scan delivery: %v", err)
		return err
	}

	return nil
}

func GetPayment(id string, data *structs.Order) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	row := db.QueryRow(`
		SELECT
			*
		FROM
			payment
		WHERE
			order_uid = $1
	`, id)

	err = row.Scan(
		&data.Payment.Transaction,
		&data.OrderUID,
		&data.Payment.RequestID,
		&data.Payment.Currency,
		&data.Payment.Provider,
		&data.Payment.Amount,
		&data.Payment.PaymentDT,
		&data.Payment.Bank,
		&data.Payment.DeliveryCost,
		&data.Payment.GoodsTotal,
		&data.Payment.CustomFee,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No delivery found for order with id: %s", id)
			return nil
		}
		log.Printf("Failed to scan delivery: %v", err)
		return err
	}

	return nil
}

func GetItems(id string, data *structs.Order) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT
			*
		FROM 
			items
		WHERE
			order_uid = $1
	`, id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return err
	}
	defer rows.Close()

	var items []structs.Item

	for rows.Next() {
		var tmpId string
		var item structs.Item
		err = rows.Scan(
			&item.ChrtID,
			&tmpId,
			&item.TrackNumber,
			&item.Price,
			&item.RID,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NMID,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			log.Printf("Error after rows iteration: %v", err)
		}
		items = append(items, item)
		if len(items) == 0 {
			log.Printf("No items found for order with id: %s", id)
			return errors.New("query is empty")
		}
	}

	data.Items = append(data.Items, items...)

	return nil
}
