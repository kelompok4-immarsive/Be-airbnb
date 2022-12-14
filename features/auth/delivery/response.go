package delivery

import "fajar/testing/features/user"

type UserResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func FromCore(data user.CoreUser) UserResponse {
	return UserResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
}
