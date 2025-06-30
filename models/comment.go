package models

import "time"

type Comment struct {
	Content   string `json:"content"`
	CreatedAt time.Time
	ID        uint `gorm:"primaryKey" json:"id"`
	Post      Post `gorm:"foreignKey:PostID"`
	PostID    uint `gorm:"column:post_id"`
	User      User `gorm:"foreignKey:UserID"`
	UserID    uint `json:"user_id"`
}
