package models

import "time"

type Comment struct {
	Content   string `json:"content"`
	CreatedAt time.Time
	ID        uint `gorm:"primaryKey" json:"id"`
	Post      Post `gorm:"foreignKey:PostID" binding:"-"`
	PostID    uint `gorm:"column:post_id" json:"postId"`
	User      User `gorm:"foreignKey:UserID" binding:"-"`
	UserID    uint `json:"userId"`
}
