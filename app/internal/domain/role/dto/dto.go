package dto

import "my_project/internal/model"

type RoleDTO struct {
	Name               string             `json:"name" form:"name" validate:"required,max=70"`
	Permissions_append []model.Permission `gorm:"many2many:roles_permissions" json:"permissions_append,omitempty"`
	Permissions_delete []model.Permission `gorm:"many2many:roles_permissions" json:"permissions_delete,omitempty"`
}
