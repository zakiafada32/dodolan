package migration

import (
	"github.com/zakiafada32/retail/modules/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
