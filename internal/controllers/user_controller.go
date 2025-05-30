package controllers

import (
	"Vault/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error hashing password"})
			return
		}

		user.Password = string(hashedPassword)

		if err := db.Create(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": "Error creating user"})
			return
		}

		c.JSON(201, gin.H{"message": "User registered successfully"})
	}
}
