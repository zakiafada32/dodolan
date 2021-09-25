package cart

import (
	"time"

	"github.com/zakiafada32/retail/modules/product"
	"github.com/zakiafada32/retail/modules/user"
)

type CartItem struct {
	UserID    string
	User      user.User
	ProductID uint32
	Product   product.Product
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
