package user

import "github.com/zakiafada32/retail/business/user"

type createNewUserRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Address  string `json:"address" validate:"required"`
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

type loginRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type updateUserRequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (req *updateUserRequestBody) convertToUpdateUserBusiness() user.UpdateUser {
	return user.UpdateUser{
		Name:    req.Name,
		Address: req.Address,
	}
}
