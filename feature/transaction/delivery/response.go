package delivery

import (
	"finance/domain"
	"time"
)

type DataResponse struct {
	ID              int       `json:"id"`
	CustomerID      int       `json:"customer_id" form:"customer_id" validate:"required"`
	ContractNumber  string    `json:"contract_number" form:"contract_number" validate:"required"`
	OTR             float64   `json:"otr" form:"otr" validate:"required"`
	AdminFee        float64   `json:"admin_fee" form:"admin_fee" validate:"required"`
	Installments    int       `json:"installments" form:"installments" validate:"required"`
	Interest        float64   `json:"interest" form:"interest" validate:"required"`
	AssetName       string    `json:"asset_name" form:"asset_name" validate:"required"`
	TransactionDate time.Time `json:"transaction_date" form:"transaction_date" validate:"required"`
}

func FromDomain(data domain.Transaction) DataResponse {
	var res DataResponse
	res.CustomerID = data.CustomerID
	res.ContractNumber = data.ContractNumber
	res.OTR = data.OTR
	res.AdminFee = data.AdminFee
	res.Installments = data.Installments
	res.Interest = data.Interest
	res.AssetName = data.AssetName
	res.TransactionDate = data.TransactionDate
	return res
}
