package data

import (
	"finance/domain"

	"gorm.io/gorm"
)

type Limit struct {
	gorm.Model
	CustomerID int     `json:"customer_id" form:"customer_id" validate:"required"`
	Tenor1     float64 `json:"tenor_1" form:"tenor_1"`
	Tenor2     float64 `json:"tenor_2" form:"tenor_2"`
	Tenor3     float64 `json:"tenor_3" form:"tenor_3"`
	Tenor6     float64 `json:"tenor_6" form:"tenor_6"`
}

func (b *Limit) ToDomain() domain.Limit {
	return domain.Limit{
		CustomerID: int(b.CustomerID),
		Tenor1:     b.Tenor1,
		Tenor2:     b.Tenor2,
		Tenor3:     b.Tenor3,
		Tenor6:     b.Tenor6,
	}
}

func ParseToArr(arr []Limit) []domain.Limit {
	var res []domain.Limit

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Limit) Limit {
	var res Limit
	res.Tenor1 = data.Tenor1
	res.Tenor2 = data.Tenor2
	res.Tenor3 = data.Tenor3
	res.Tenor6 = data.Tenor6
	return res
}
