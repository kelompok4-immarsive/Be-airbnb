package delivery

import "fajar/testing/features/user"

type RequestUser struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

func RequestUserToCoreUser(data RequestUser) user.CoreUser {
	return user.CoreUser{
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
		Status:   data.Status,
	}
}
