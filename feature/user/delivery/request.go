package delivery

import (
	"finance/domain"
)

type InsertFormat struct {
	Username string `gorm:"unique" json:"username" form:"username" validate:"required"`
	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"default:user"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		UserName: i.Username,
		Email:    i.Email,
		Password: i.Password,
		FullName: i.FullName,
		Role:     i.Role,
	}
}

type LoginFormat struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (lf *LoginFormat) LoginToModel() domain.User {
	return domain.User{
		UserName: lf.UserName,
		Password: lf.Password,
	}
}
