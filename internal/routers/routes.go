package routers

import (
	"Vault/internal/controllers"
	"Vault/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	public := router.Group("/public")
	{
		public.POST("register", controllers.RegisterUser(db))
		public.POST("login", controllers.LoginUser(db))
	}
	private := router.Group("/private")
	private.Use(middlewares.AuthMiddleware())
	{
		private.GET("sexo", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "sexo",
			})
		})
	}
	return router
}
