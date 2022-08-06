package model

import (
	"time"
)

type Permission struct {
	ID        int       `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Roles     []Role    `gorm:"many2many:roles_permissions" json:"roles,omitempty"`
}
