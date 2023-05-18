package domain

import (
	"github.com/labstack/echo/v4"
)

type Limit struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Tenor1     float64 `json:"tenor_1"`
	Tenor2     float64 `json:"tenor_2"`
	Tenor3     float64 `json:"tenor_3"`
	Tenor6     float64 `json:"tenor_6"`
}

type LimitHandler interface {
	InsertLimit() echo.HandlerFunc
	UpdateLimit() echo.HandlerFunc
	DeleteLimit() echo.HandlerFunc
	GetLimitID() echo.HandlerFunc
}

type LimitUseCase interface {
	AddLimit(newData Limit) (Limit, error)
	UpLimit(IDLimit int, updateData Limit) (Limit, error)
	DelLimit(IDLimit int) (bool, error)
	GetSpecificLimit(LimitID int) ([]Limit, error)
}

type LimitData interface {
	Insert(insertLimit Limit) Limit
	Update(IDLimit int, updatedLimit Limit) Limit
	Delete(IDLimit int) error
	GetLimitID(LimitID int) []Limit
}
