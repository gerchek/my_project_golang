package dto

import "my_project/internal/model"

type AdminDTO struct {
	Username             string  `json:"username" form:"username" validate:"required"`
	FirstName            string  `json:"first_name" form:"first_name" validate:"required"`
	LastName             string  `json:"last_name" form:"last_name" validate:"required"`
	Password             *string `json:"password" form:"password" validate:"omitempty,min=8"`
	PasswordConfirmation *string `json:"password_confirmation" form:"password_confirmation" validate:"omitempty,min=8,eqfield=Password"`
	// Roles                []model.Role `json:"roles,omitempty" form:"roles,omitempty"`

	Roles_append []model.Role `gorm:"many2many:roles_permissions" json:"roles_append,omitempty"`
	Roles_delete []model.Role `gorm:"many2many:roles_permissions" json:"roles_delete,omitempty"`
	// Permissions          []model.Permission `json:"permissions,omitempty" form:"permissions,omitempty"`
}

type AdminLoginDTO struct {
	Username string `json:"username" form:"username" validate:"required,min=3"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
}
