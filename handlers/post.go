package handlers

import (
	"blog-system/config"
	"blog-system/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	username, _ := c.Get("claims")
	user := username.(models.User)
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post.UserID = user.ID
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "无法创建帖子"})
		return
	}
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.DB.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "无法获取帖子"})
		return
	}
	c.JSON(200, posts)
}

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := config.DB.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "帖子不存在"})
		return
	}
	c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "帖子不存在"})
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "无法更新帖子"})
		return
	}
}

func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "帖子不存在"})
		return
	}
	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "无法删除帖子"})
		return
	}
}

func GetUserPosts(c *gin.Context) {
	var posts []models.Post
	username, _ := c.Get("claims")
	user := username.(models.User)
	if err := config.DB.Preload("User").Where("user_id = ?", user.ID).Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "无法获取帖子"})
		return
	}
	c.JSON(200, posts)
}
