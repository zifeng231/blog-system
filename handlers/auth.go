package handlers

import (
	"blog-system/config"
	"blog-system/models"
	"blog-system/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册用户
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//检查用户名和邮箱是否存在
	var existUser models.User
	if config.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existUser).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户名或邮箱已存在",
		})
		return
	}
	//加密密码
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})

}

// 登录
func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}
	// 验证密码
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}
	// 生成 JWT
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成 token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
