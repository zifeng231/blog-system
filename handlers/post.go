package handlers

import (
	"blog-system/config"
	"blog-system/models"
	"blog-system/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	username, _ := c.Get("claims")
	claims := username.(*utils.Claims)
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post.UserID = claims.User.ID
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "无法创建帖子"})
		return
	}
	c.JSON(200, gin.H{"message": "帖子创建成功"})
}

func GetPosts(c *gin.Context) {
	var query struct {
		Title    string `json:"title" `
		Content  string `json:"content"`
		UserID   string `json:"userId"`
		PageNum  int    `json:"pageNum" binding:"omitempty,min=1"`
		PageSize int    `json:"pageSize" binding:"omitempty,min=1,max=100"`
	}
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(400, gin.H{"error": utils.TranslateValidationErrors(err)})
		return
	}
	//设置默认值
	if query.PageNum == 0 {
		query.PageNum = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 10
	}
	//获取当前的用户信息
	claims, _ := c.Get("claims")
	claim := claims.(*utils.Claims)
	userId := claim.User.ID
	if query.UserID != "" {
		userIdInt, err := strconv.ParseUint(query.UserID, 10, 64)
		if err != nil || uint(userIdInt) != userId {
			c.JSON(http.StatusForbidden, gin.H{"error": "无权限访问"})
		}
	}
	//构建查询
	db := config.DB.Model(&models.Post{}).Preload("User")
	//添加查询条件
	if query.Title != "" {
		db.Where("title LIKE ?", "%"+query.Title+"%")
	}
	if query.Content != "" {
		db.Where("content LIKE ?", "%"+query.Content+"%")
	}
	db.Where("user_id = ?", userId)
	//查询总数
	var total int64
	db.Count(&total)
	//分页查询
	var posts []models.Post
	offset := (query.PageNum - 1) * query.PageSize
	db.Offset(offset).Limit(query.PageSize).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"items": posts,
	})
}

func GetPost(c *gin.Context) {
	var post models.Post
	//从路劲中获取参数
	id := c.Param("id")
	if err := config.DB.Preload("User").First(&post, id).Error; err != nil {
		//c.JSON(404, gin.H{"error": "帖子不存在"})
		c.Error(utils.ErrBadRequest)
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
	c.JSON(200, gin.H{"message": "成功"})
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
