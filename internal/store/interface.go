package store

import (
	"github.com/jmoiron/sqlx"
)

var _ StoreManager = (*StoreManagement)(nil)

type Id = int64
type Name = string
type Size = string
type Code = string
type Quantity = int
type Available = bool

type Product struct {
	ID       Id       `json:"id"`
	Name     Name     `json:"name"`
	Size     Size     `json:"size"`
	Code     Code     `json:"code"`
	Quantity Quantity `json:"quantity"`
	StoreID  Id       `json:"store_id"`
}

type Store struct {
	ID        Id        `json:"id" db:"id"`
	Name      Name      `json:"name" db:"name"`
	Available Available `json:"is_available" db:"is_available"`
}

type StoreManagement struct {
	Db *sqlx.DB
}

type StoreManager interface {
	CreateStore(name Name, isAvailable Available) (Id, error)
	CreateProduct(name Name, size Size, code Code, quantity Quantity, storeID Id) (Id, error)
	DeleteProduct(id Id) error
	ReserveProducts(productCodes []Code) error
	ReleaseProducts(productCodes []Code) error
	GetRemainingProducts(storeID Id) ([]Product, error)
}
