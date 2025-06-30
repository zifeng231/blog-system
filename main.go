package main

import (
	"blog-system/config"
	"blog-system/handlers"
	"blog-system/middleware"
	"blog-system/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db := config.ConnectDatabase()
	// 2. 自动迁移模型（创建表）
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "欢迎来到个人博客系统",
		})
	})
	//注册和登录接口不需要认证
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	//需要认证的接口
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("claims")
			c.JSON(http.StatusOK, gin.H{"message": "欢迎回来", "username": username})
		})

		protected.POST("/posts", handlers.CreatePost)
		protected.GET("/posts", handlers.GetPosts)
		protected.GET("/posts/:id", handlers.GetPost)
		protected.PUT("/posts/:id", handlers.UpdatePost)
		protected.DELETE("/posts/:id", handlers.DeletePost)
		protected.POST("/posts/user", handlers.GetUserPosts)
		protected.GET("/comments/:id", handlers.CreateComment)
		protected.GET("/comments-user/:id", handlers.GetComments)

	}
	r.Run(":8080")
}
