package dto

type CommentDTO struct {
	CommentText string `json:"comment_text" form:"comment_text" validate:"required,max=255"`
	AdminID     uint   `json:"admin_id" form:"admin_id" validate:"required"`
	ProductID   uint   `json:"product_id" form:"product_id" validate:"required"`
}
