package dto

type PermissionDTO struct {
	Name string `json:"name" form:"name" validate:"required,max=70"`
}
