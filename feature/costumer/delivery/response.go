package delivery

import (
	"database/sql"
	"finance/domain"
)

type DataResponse struct {
	ID           int          `json:"id"`
	NIK          string       `json:"nik" gorm:"unique" form:"nik" validate:"required"`
	FullName     string       `json:"full_name"  form:"full_name" validate:"required"`
	LegalName    string       `json:"legal_name" form:"legal_name"`
	PlaceOfBirth string       `json:"place_of_birth" form:"place_of_birth"`
	DateOfBirth  sql.NullTime `json:"date_of_birth" form:"date_of_birth"`
	Salary       float64      `json:"salary" form:"salary"`
	KTPPhoto     string       `json:"ktp_photo" form:"ktp_photo"`
	SelfiePhoto  string       `json:"selfie_photo" form:"selfie_photo"`
}

func FromDomain(data domain.Costumer) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.NIK = data.NIK
	res.FullName = data.FullName
	res.LegalName = data.LegalName
	res.PlaceOfBirth = data.PlaceOfBirth
	res.DateOfBirth = data.DateOfBirth
	res.Salary = data.Salary
	res.KTPPhoto = data.KTPPhoto
	return res
}
