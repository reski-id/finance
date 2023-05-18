package data

import (
	"finance/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" form:"username" validate:"required"`
	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"default:user"`
}

func (u *User) ToDomain() domain.User {
	return domain.User{
		ID:        int(u.ID),
		UserName:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FullName:  u.FullName,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) ToDomain2() domain.User {
	return domain.User{
		ID:        int(u.ID),
		UserName:  u.Username,
		Email:     u.Email,
		FullName:  u.FullName,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}

	return res
}

func ParseToArr2(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToDomain2())
	}

	return res
}

func ToLocal(data domain.User) User {
	var res User
	res.Username = data.UserName
	res.Email = data.Email
	res.Password = data.Password
	res.FullName = data.FullName
	res.Role = data.Role
	return res
}
