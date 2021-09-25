package courier

import (
	"github.com/zakiafada32/retail/business/courier"
)

type createNewCourierProviderRequestBody struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (req *createNewCourierProviderRequestBody) convertToCourierProviderBusiness() courier.CourierProvider {
	return courier.CourierProvider{
		Name:        req.Name,
		Description: req.Description,
	}
}
