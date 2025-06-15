package controllers

import (
	"Vault/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error(), "details": "Invalid input"})
			return
		}
		if !models.CheckUserExists(db, user.Email) {
			c.JSON(400, gin.H{"error": "User already exists", "details": "Email is already registered"})
			return
		}
		if models.CreateUser(db, user, c) != nil {
			c.JSON(500, gin.H{"error": "Failed to create user", "details": "Internal server error"})
			return
		}
		c.JSON(201, gin.H{"message": "User registered successfully"})
	}
}

func LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error(), "details": "Invalid input"})
			return
		}
		ex := models.CheckUserExists(db, user.Email)
		if ex {
			c.JSON(400, gin.H{"error": "User does not exist", "details": "Email is not registered"})
			return
		}
		cre := models.UserMachtchesPassword(db, user.Email, user.Password)
		if !cre {
			c.JSON(400, gin.H{"error": "Invalid credentials", "details": "Email or password is incorrect"})
			return
		} else {
			c.JSON(200, gin.H{"message": "Login successful", "user": user})
		}
	}
}
