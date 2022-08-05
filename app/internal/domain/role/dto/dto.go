package dto

type RoleDTO struct {
	Name string `json:"name" form:"name" validate:"required,max=70"`
}
