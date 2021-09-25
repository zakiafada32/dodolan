package user

import (
	"time"

	"github.com/zakiafada32/retail/business/user"
)

type UserResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func convertToUserResponse(user user.User) UserResponse {
	return UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
