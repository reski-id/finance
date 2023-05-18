package delivery

import "finance/domain"

type DataResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Role     string `json:"role"`
}

func FromDomain(data domain.User) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.Username = data.UserName
	res.Email = data.Email
	res.Password = data.Password
	res.Role = data.Role
	return res
}
