package models

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique" binding:"email,required"`
	Password string
	Phone    string `gorm:"unique" binding:"required"`
}

func CheckUserExists(db *gorm.DB, email string) bool {
	var user User
	db.Where("email = ?", email).First(&user)
	return user.ID == 0
}

func CreateUser(db *gorm.DB, user User, c *gin.Context) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "details": "Failed to hash password"})
		return err
	}
	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "details": "Failed to create user"})
		return err
	}
	c.JSON(201, gin.H{"message": "User registered successfully"})
	return nil
}
func UserMachtchesPassword(db *gorm.DB, email, password string) bool {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true
}
