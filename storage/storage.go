package storage

import (
	"github.com/Dostonlv/gorm-auth/models"
	"github.com/Dostonlv/gorm-auth/storage/postgres"
)

func SyncDatabase() {
	postgres.DB.AutoMigrate(&models.User{})
}
