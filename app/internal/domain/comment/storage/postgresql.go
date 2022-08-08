package storage

import (
	"my_project/internal/model"

	"gorm.io/gorm"
)

type CommentStorage interface {
	Create(role *model.Comment) error
}

type commentStorage struct {
	client *gorm.DB
}

func NewCommentStorage(client *gorm.DB) CommentStorage {
	return &commentStorage{
		client: client,
	}
}

func (rs *commentStorage) Create(comment *model.Comment) error {
	if err := rs.client.Create(comment).Error; err != nil {
		return err
	}
	return nil
}
