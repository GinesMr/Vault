package routers

import (
	"Vault/internal/controllers"
	"Vault/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	public := router.Group("/public")
	{
		public.POST("register", controllers.RegisterUser(db))
		public.POST("login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "register success",
			})
		})
	}
	private := router.Group("/private")
	private.Use(middleware.AuthMiddleware())
	{
		private.GET("sexo", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "sexo",
			})
		})
	}
	return router
}
