package user

import "github.com/zakiafada32/retail/business/user"

type CreateNewUserRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Name     string `json:"name" validate:"required"`
}

func (b *CreateNewUserRequestBody) convertToUserBusiness() user.User {
	return user.User{
		Name:     b.Name,
		Email:    b.Email,
		Password: b.Password,
	}
}
