package model

import (
	"time"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	AdminID   int       `json:"admin_id" gorm:"column:admin_id"`
	Comments  []Comment `gorm:"ForeignKey:ProductID"`
	// Admin     Admin
}
