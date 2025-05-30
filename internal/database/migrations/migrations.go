package migrations

import (
	"fmt"
	"gorm.io/gorm"
)

func SetUpMigration(db *gorm.DB) error {
	if err := user_Migration(db); err != nil {
		return fmt.Errorf("failed to migrate user: %w", err)
	}
	return nil
}
