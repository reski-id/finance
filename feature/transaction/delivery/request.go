package delivery

import (
	"finance/domain"
	"time"
)

type InsertRequest struct {
	CustomerID      int       `json:"customer_id" form:"customer_id" validate:"required"`
	ContractNumber  string    `json:"contract_number" form:"contract_number" validate:"required"`
	OTR             float64   `json:"otr" form:"otr" validate:"required"`
	AdminFee        float64   `json:"admin_fee" form:"admin_fee" validate:"required"`
	Installments    int       `json:"installments" form:"installments" validate:"required"`
	Interest        float64   `json:"interest" form:"interest" validate:"required"`
	AssetName       string    `json:"asset_name" form:"asset_name" validate:"required"`
	TransactionDate time.Time `json:"transaction_date" form:"transaction_date" validate:"required"`
}

func (ni *InsertRequest) ToDomain() domain.Transaction {
	return domain.Transaction{
		CustomerID:      ni.CustomerID,
		ContractNumber:  ni.ContractNumber,
		OTR:             ni.OTR,
		AdminFee:        ni.AdminFee,
		Installments:    ni.Installments,
		Interest:        ni.Interest,
		AssetName:       ni.AssetName,
		TransactionDate: ni.TransactionDate,
	}
}
