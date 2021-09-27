package cart

import (
	"time"

	"github.com/zakiafada32/retail/modules/category"
	"github.com/zakiafada32/retail/modules/user"
)

type CartItem struct {
	UserID    string
	User      user.User
	ProductID uint32
	Product   category.Product
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
