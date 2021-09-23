package user

import "github.com/zakiafada32/retail/business/user"

type CreateNewUserRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

func (c *CreateNewUserRequestBody) convertToUserBusiness() user.User {
	return user.User{
		Name:     c.Name,
		Email:    c.Email,
		Password: c.Password,
	}
}

type LoginRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}
