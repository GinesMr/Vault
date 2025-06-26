package main

import (
	"Vault/config"
	"Vault/internal/database"
	"Vault/internal/database/migrations"
	"Vault/internal/routers"
)

func main() {
	config.LoadEnvConfig()
	var db = database.DbConex()
	migrations.SetUpMigration(db)
	router := routers.LoadRoutes(db)
	router.Run("0.0.0.0:8080")
}
