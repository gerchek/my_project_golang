package model

import (
	"time"
)

type Role struct {
	ID        int       `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Admins    []Admin   `gorm:"many2many:admins_roles" json:"admins,omitempty"`
}
