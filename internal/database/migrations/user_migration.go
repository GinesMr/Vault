package migrations

import (
	"Vault/internal/models"
	"gorm.io/gorm"
	"log"
)

func user_Migration(db *gorm.DB) error {
	err := db.AutoMigrate(models.User{})

	log.Println("Database migrated")
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
