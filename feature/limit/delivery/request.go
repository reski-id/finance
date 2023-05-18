package delivery

import (
	"finance/domain"
)

type InsertRequest struct {
	CustomerID int     `json:"customer_id" form:"customer_id" validate:"required"`
	Tenor1     float64 `json:"tenor_1" form:"tenor_1"`
	Tenor2     float64 `json:"tenor_2" form:"tenor_2"`
	Tenor3     float64 `json:"tenor_3" form:"tenor_3"`
	Tenor6     float64 `json:"tenor_6" form:"tenor_6"`
}

func (ni *InsertRequest) ToDomain() domain.Limit {
	return domain.Limit{
		CustomerID: ni.CustomerID,
		Tenor1:     ni.Tenor1,
		Tenor2:     ni.Tenor2,
		Tenor3:     ni.Tenor3,
		Tenor6:     ni.Tenor6,
	}
}
