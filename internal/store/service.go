package store

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func NewStoreManagement(db *sqlx.DB) *StoreManagement {
	return &StoreManagement{
		Db: db,
	}
}

func (s *StoreManagement) CreateStore(name Name, isAvailable Available) (Id, error) {
	const query = "INSERT INTO stores(name, is_available) VALUES($1, $2) RETURNING id"

	var id Id
	err := s.Db.QueryRowx(query, name, isAvailable).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("Cannot make CreateStore query: %w", err)
	}

	return id, nil
}

func (s *StoreManagement) CreateProduct(name Name, size Size, code Code, quantity Quantity, storeID Id) (Id, error) {
	const query = "INSERT INTO products(name, size, code, quantity, store_id) VALUES($1, $2, $3, $4, $5) RETURNING id"

	var id Id
	err := s.Db.QueryRow(query, name, size, code, quantity, storeID).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("Cannot make CreateProduct query: %w", err)
	}

	return id, nil
}

func (s *StoreManagement) DeleteProduct(id Id) error {
	_, err := s.Db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoreManagement) ReserveProducts(productCodes []Code) error {
	if len(productCodes) == 0 {
		return errors.New("There are no product code")
	}

	tx, err := s.Db.Begin()
	if err != nil {
		return fmt.Errorf("Cannot begin transaction: %w", err)
	}

	for iCode := range productCodes {
		row := tx.QueryRow("SELECT id, name, size, code, quantity FROM products WHERE code = $1 FOR UPDATE", productCodes[iCode])

		var p Product
		err := row.Scan(&p.ID, &p.Name, &p.Size, &p.Code, &p.Quantity)
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.Printf("Cannot get the result: %v. Cannot make rollback: %v\n", err, rollBackErr)
			}
			return fmt.Errorf("Cannot scan the result: %w", err)
		}

		if p.Quantity < 1 {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.Printf("Cannot make rollback: %v\n", rollBackErr)
			}
			return errors.New("Product is out of stock")
		}

		_, err = tx.Exec("UPDATE products SET quantity = quantity - 1 WHERE id = $1", p.ID)
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.Printf("Cannot execute query: %v. Cannot make rollback: %v\n", err, rollBackErr)
			}
			return fmt.Errorf("Cannot execute query: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Printf("Cannot make commit: %v. Cannot make rollback: %v\n", err, rollBackErr)
		}
		return fmt.Errorf("Cannot commit: %w", err)
	}

	return nil
}

func (s *StoreManagement) ReleaseProducts(productCodes []Code) error {
	if len(productCodes) == 0 {
		return errors.New("There are no product code")
	}

	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}

	for iCode := range productCodes {
		var p Product
		err := tx.QueryRow("SELECT id, name, size, code, quantity FROM products WHERE code = $1", productCodes[iCode]).Scan(&p.ID, &p.Name, &p.Size, &p.Code, &p.Quantity)
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.Printf("Cannot get the result: %v. Cannot make rollback: %v\n", err, rollBackErr)
			}
			return fmt.Errorf("Cannot scan the result: %w", err)
		}

		_, err = tx.Exec("UPDATE products SET quantity = quantity + 1 WHERE code = $1", p.Code)
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.Printf("Cannot execute query: %v. Cannot make rollback: %v\n", err, rollBackErr)
			}
			return fmt.Errorf("Cannot execute query: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Printf("Cannot make commit: %v. Cannot make rollback: %v\n", err, rollBackErr)
		}
		return fmt.Errorf("Cannot commit: %w", err)

	}

	return nil
}

func (s *StoreManagement) GetRemainingProducts(storeID Id) ([]Product, error) {
	rows, err := s.Db.Query("SELECT code, quantity FROM products WHERE store_id = $1", storeID)
	if err != nil {
		return nil, fmt.Errorf("Cannot execute query: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Code, &p.Quantity); err != nil {
			return nil, fmt.Errorf("Cannot scan the result: %w", err)
		}
		p.StoreID = storeID
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Got error in row of the result: %w", err)
	}

	return products, nil
}
