package domain

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type Costumer struct {
	ID           int
	NIK          string
	FullName     string
	LegalName    string
	PlaceOfBirth string
	DateOfBirth  sql.NullTime
	Salary       float64
	KTPPhoto     string
	SelfiePhoto  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Transaction struct {
	ContractNumber  string    `json:"contract_number"`
	OTR             float64   `json:"otr"`
	AdminFee        float64   `json:"admin_fee"`
	Installments    int       `json:"installments"`
	Interest        float64   `json:"interest"`
	AssetName       string    `json:"asset_name"`
	CustomerID      int       `json:"customer_id"`
	TransactionID   string    `json:"transaction_id"`
	TransactionDate time.Time `json:"transaction_date"`
}

type CostumerHandler interface {
	InsertCostumer() echo.HandlerFunc
	UpdateCostumer() echo.HandlerFunc
	DeleteCostumer() echo.HandlerFunc
	GetAllCostumer() echo.HandlerFunc
	GetCostumerID() echo.HandlerFunc
}

type CostumerUseCase interface {
	AddCostumer(newData Costumer) (Costumer, error)
	UpCostumer(IDCostumer int, updateData Costumer) (Costumer, error)
	DelCostumer(IDCostumer int) (bool, error)
	GetAllData() ([]Costumer, error)
	GetSpecificCostumer(CostumerID int) ([]Costumer, error)
}

type CostumerData interface {
	Insert(insertCostumer Costumer) Costumer
	Update(IDCostumer int, updatedCostumer Costumer) Costumer
	Delete(IDCostumer int) error
	GetAll() []Costumer
	GetCostumerID(CostumerID int) []Costumer
}
