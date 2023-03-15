package dto

type ProductDTO struct {
	Name    string `json:"name" form:"name" validate:"required,max=255"`
	AdminID uint   `json:"admin_id" form:"admin_id" validate:"required"`
}
