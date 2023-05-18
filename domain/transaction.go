package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID              int       `json:"id" form:"id"`
	CustomerID      int       `json:"customer_id" form:"customer_id" validate:"required"`
	ContractNumber  string    `json:"contract_number" form:"contract_number" validate:"required"`
	OTR             float64   `json:"otr" form:"otr" validate:"required"`
	AdminFee        float64   `json:"admin_fee" form:"admin_fee" validate:"required"`
	Installments    int       `json:"installments" form:"installments" validate:"required"`
	Interest        float64   `json:"interest" form:"interest" validate:"required"`
	AssetName       string    `json:"asset_name" form:"asset_name" validate:"required"`
	TransactionDate time.Time `json:"transaction_date" form:"transaction_date" validate:"required"`
}

type TransactionHandler interface {
	InsertTransaction() echo.HandlerFunc
	GetTransactionID() echo.HandlerFunc
}

type TransactionUseCase interface {
	AddTransaction(newData Transaction) (Transaction, error)
	GetSpecificTransaction(TransactionID int) ([]Transaction, error)
}

type TransactionData interface {
	Insert(insertTransaction Transaction) Transaction
	GetTransactionID(TransactionID int) []Transaction
}
