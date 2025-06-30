package handlers

import (
	"blog-system/config"
	"blog-system/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// 创建评论
func CreateComment(c *gin.Context) {
	postID := c.Param("id")
	var req models.Comment

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 检查关联的文章是否存在
	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err := config.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(http.StatusCreated, req)
}

// 获取文章的所有评论
func GetComments(c *gin.Context) {
	postID := c.Param("id")
	// 检查关联的文章是否存在
	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var comments []models.Comment
	if err := config.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
