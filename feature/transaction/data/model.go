package data

import (
	"finance/domain"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CustomerID      int       `json:"customer_id" form:"customer_id" validate:"required"`
	ContractNumber  string    `json:"contract_number" form:"contract_number" validate:"required"`
	OTR             float64   `json:"otr" form:"otr" validate:"required"`
	AdminFee        float64   `json:"admin_fee" form:"admin_fee" validate:"required"`
	Installments    int       `json:"installments" form:"installments" validate:"required"`
	Interest        float64   `json:"interest" form:"interest" validate:"required"`
	AssetName       string    `json:"asset_name" form:"asset_name" validate:"required"`
	TransactionDate time.Time `json:"transaction_date" form:"transaction_date" validate:"required"`
}

func (b *Transaction) ToDomain() domain.Transaction {
	return domain.Transaction{
		CustomerID:      int(b.CustomerID),
		ContractNumber:  b.ContractNumber,
		OTR:             b.OTR,
		AdminFee:        b.AdminFee,
		Installments:    b.Installments,
		Interest:        b.Interest,
		AssetName:       b.AssetName,
		TransactionDate: b.TransactionDate,
	}
}

func ParseToArr(arr []Transaction) []domain.Transaction {
	var res []domain.Transaction

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Transaction) Transaction {
	var res Transaction
	res.CustomerID = data.CustomerID
	res.ContractNumber = data.ContractNumber
	res.OTR = data.OTR
	res.AdminFee = data.AdminFee
	res.Installments = data.Installments
	res.Interest = data.Interest
	res.AssetName = data.AssetName
	return res
}
