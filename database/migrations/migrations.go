package migrations

import (
	"github.com/pedroteixeirabisognin/swapi-go/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.Planet{})
}
