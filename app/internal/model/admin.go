package model

import "time"

type Admin struct {
	ID           uint64    `json:"id"`
	Username     string    `json:"username"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Password     string    `json:"-"`
	AccessToken  string    `gorm:"-" json:"access_token,omitempty"`
	RefreshToken string    `gorm:"-" json:"refresh_token,omitempty"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	Roles        []Role    `gorm:"many2many:admins_roles" json:"roles,omitempty"`
	MyProducts   []Product `gorm:"ForeignKey:AdminID"`
	MyComments   []Comment `gorm:"ForeignKey:AdminID"`
	// Permissions  []Permission `gorm:"many2many:admin_permissions" json:"permissions,omitempty"`
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
