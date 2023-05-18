package delivery

import (
	"database/sql"
	"finance/domain"
)

type InsertRequest struct {
	NIK          string       `json:"nik" gorm:"unique" form:"nik" validate:"required"`
	FullName     string       `json:"full_name"  form:"full_name" validate:"required"`
	LegalName    string       `json:"legal_name" form:"legal_name"`
	PlaceOfBirth string       `json:"place_of_birth" form:"place_of_birth"`
	DateOfBirth  sql.NullTime `json:"date_of_birth" form:"date_of_birth"`
	Salary       int          `json:"salary" form:"salary"`
	KTPPhoto     string       `json:"ktp_photo" form:"ktp_photo"`
	SelfiePhoto  string       `json:"selfie_photo" form:"selfie_photo"`
}

func (ni *InsertRequest) ToDomain() domain.Costumer {
	return domain.Costumer{
		NIK:          ni.NIK,
		FullName:     ni.FullName,
		LegalName:    ni.LegalName,
		PlaceOfBirth: ni.PlaceOfBirth,
		DateOfBirth:  ni.DateOfBirth,
		Salary:       float64(ni.Salary),
		KTPPhoto:     ni.KTPPhoto,
		SelfiePhoto:  ni.SelfiePhoto,
	}
}
