package service

import (
	"errors"
	"my_project/internal/domain/comment/dto"
	"my_project/internal/domain/comment/storage"
	"my_project/internal/model"

	"github.com/go-redis/redis/v8"
)

type CommentService interface {
	Create(commentDTO *dto.CommentDTO) error
}

type commentService struct {
	storage     storage.CommentStorage
	redisClient *redis.Client
}

func NewCommentService(storage storage.CommentStorage, redisClient *redis.Client) CommentService {
	return &commentService{
		storage:     storage,
		redisClient: redisClient,
	}
}

func (s *commentService) Create(commentDTO *dto.CommentDTO) error {
	admin := &model.Comment{
		CommentText: commentDTO.CommentText,
		AdminID:     int(commentDTO.AdminID),
		ProductID:   int(commentDTO.ProductID),
	}

	err := s.storage.Create(admin)
	if err != nil {
		return err
	}
	if err != nil {
		return errors.New("admin created but there was problem when updating order number")
	}
	return nil
}
