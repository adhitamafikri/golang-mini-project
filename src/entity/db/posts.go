package db

import "time"

type Posts struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	Title      string     `gorm:"column:title" json:"title"`
	Body       string     `gorm:"column:body" json:"body"`
	PictureURL string     `gorm:"column:picture_url" json:"picture_url"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	User       *User      `gorm:"auto_preload" gorm:"foreignkey:UserID"`
	UserID     uint       `gorm:"column:user_id"`
}

func (Posts) TableName() string {
	return "posts"
}
