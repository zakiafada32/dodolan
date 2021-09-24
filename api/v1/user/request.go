package user

import "github.com/zakiafada32/retail/business/user"

type createNewUserRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Address  string `json:"address"`
	IsAdmin  bool   `json:"is_admin"`
}

func (req *createNewUserRequestBody) convertToUserBusiness() user.User {
	return user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		IsAdmin:  req.IsAdmin,
	}
}

type LoginRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
