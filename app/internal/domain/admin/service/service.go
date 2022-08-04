package service

import (
	"my_project/internal/domain/admin/storage"

	"github.com/go-redis/redis/v8"
)

type AdminService interface {
	All() string
}

type adminService struct {
	storage     storage.AdminStorage
	redisClient *redis.Client
}

func NewAdminService(storage storage.AdminStorage, redisClient *redis.Client) AdminService {
	return &adminService{
		storage:     storage,
		redisClient: redisClient,
	}
}

func (s *adminService) All() string {
	return s.storage.All()
}
