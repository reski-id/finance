package delivery

import (
	"finance/domain"
)

type DataResponse struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Tenor1     float64 `json:"tenor_1"`
	Tenor2     float64 `json:"tenor_2"`
	Tenor3     float64 `json:"tenor_3"`
	Tenor6     float64 `json:"tenor_6"`
}

func FromDomain(data domain.Limit) DataResponse {
	var res DataResponse
	res.CustomerID = data.CustomerID
	res.Tenor1 = data.Tenor1
	res.Tenor2 = data.Tenor2
	res.Tenor3 = data.Tenor3
	res.Tenor6 = data.Tenor6
	return res
}
