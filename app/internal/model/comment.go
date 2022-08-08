package model

import (
	"time"
)

type Comment struct {
	ID          int       `json:"id"`
	CommentText string    `json:"comment_text,omitempty"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	AdminID     int       `json:"admin_id" gorm:"column:admin_id"`
	ProductID   int       `json:"product_id" gorm:"column:product_id"`
}
