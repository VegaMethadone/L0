package query

import (
	"L0/internal/bd/connection"
	"fmt"
	"log"
)

func CleanData(id string) {
	deleteItemsData(id)
	deletePaymentData(id)
	deleteDeliveryData(id)
	deleteOrderData(id)
}

func deleteOrderData(id string) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	query := `
		DELETE
		FROM 
			orders
		WHERE
			order_uid = $1
	`

	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return err
	}
	// Проверка количества удаленных строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted for order_uid: %s", id)
	}

	log.Printf("Successfully deleted %d row(s) with order_uid: %s", rowsAffected, id)

	return nil
}

func deleteDeliveryData(id string) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	query := `
		DELETE
		FROM 
			delivery
		WHERE
			order_uid = $1
	`

	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return err
	}
	// Проверка количества удаленных строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted for order_uid: %s", id)
	}

	log.Printf("Successfully deleted %d row(s) with order_uid: %s", rowsAffected, id)

	return nil
}

func deletePaymentData(id string) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	query := `
		DELETE
		FROM 
			payment
		WHERE
			order_uid = $1
	`

	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return err
	}
	// Проверка количества удаленных строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted for order_uid: %s", id)
	}

	log.Printf("Successfully deleted %d row(s) with order_uid: %s", rowsAffected, id)

	return nil
}

func deleteItemsData(id string) error {
	db, err := connection.DB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err
	}
	defer db.Close()

	query := `
		DELETE
		FROM 
			items
		WHERE
			order_uid = $1
	`

	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return err
	}
	// Проверка количества удаленных строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted for order_uid: %s", id)
	}

	log.Printf("Successfully deleted %d row(s) with order_uid: %s", rowsAffected, id)

	return nil
}
