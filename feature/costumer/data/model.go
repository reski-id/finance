package data

import (
	"database/sql"
	"finance/domain"

	"gorm.io/gorm"
)

type Costumer struct {
	gorm.Model
	NIK          string       `json:"nik" gorm:"unique" form:"nik" validate:"required"`
	FullName     string       `json:"full_name"  form:"full_name" validate:"required"`
	LegalName    string       `json:"legal_name" form:"legal_name"`
	PlaceOfBirth string       `json:"place_of_birth" form:"place_of_birth"`
	DateOfBirth  sql.NullTime `json:"date_of_birth" form:"date_of_birth"`
	Salary       float64      `json:"salary" form:"salary"`
	KTPPhoto     string       `json:"ktp_photo" form:"ktp_photo"`
	SelfiePhoto  string       `json:"selfie_photo" form:"selfie_photo"`
}

func (b *Costumer) ToDomain() domain.Costumer {
	return domain.Costumer{
		ID:           int(b.ID),
		NIK:          b.NIK,
		FullName:     b.FullName,
		LegalName:    b.LegalName,
		PlaceOfBirth: b.PlaceOfBirth,
		DateOfBirth:  b.DateOfBirth,
		Salary:       b.Salary,
		KTPPhoto:     b.KTPPhoto,
		SelfiePhoto:  b.SelfiePhoto,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}
}

func ParseToArr(arr []Costumer) []domain.Costumer {
	var res []domain.Costumer

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Costumer) Costumer {
	var res Costumer
	res.ID = uint(data.ID)
	res.NIK = data.NIK
	res.FullName = data.FullName
	res.LegalName = data.LegalName
	res.PlaceOfBirth = data.PlaceOfBirth
	res.DateOfBirth = data.DateOfBirth
	res.Salary = data.Salary
	res.KTPPhoto = data.KTPPhoto
	res.SelfiePhoto = data.SelfiePhoto
	return res
}
