package service

import (
	"my_project/internal/domain/permission/dto"
	"my_project/internal/domain/permission/storage"
	"my_project/internal/model"

	"github.com/go-redis/redis/v8"
)

type PermissionService interface {
	All() []*model.Permission
	Create(permissionDTO *dto.PermissionDTO) error
	Update(permissionDTO *dto.PermissionDTO, id int) error
	Delete(id int) error
}

type permissionService struct {
	storage     storage.PermissionStorage
	redisClient *redis.Client
}

func NewPermissionService(storage storage.PermissionStorage, redisClient *redis.Client) PermissionService {
	return &permissionService{
		storage:     storage,
		redisClient: redisClient,
	}
}

func (s *permissionService) All() []*model.Permission {
	// a := s.storage.Test()
	return s.storage.All()
}

func (s *permissionService) Create(permissionDTO *dto.PermissionDTO) error {
	permission := &model.Permission{
		Name: permissionDTO.Name,
	}
	err := s.storage.Create(permission)
	if err != nil {
		return err
	}
	return nil
}

func (s *permissionService) Update(permissionDTO *dto.PermissionDTO, id int) error {
	var oldPermission model.Permission
	err := s.storage.FindByID(&oldPermission, id)
	if err != nil {
		return err
	}
	oldPermission.Name = permissionDTO.Name
	err = s.storage.Update(&oldPermission)
	if err != nil {
		return err
	}
	return nil
}

func (s *permissionService) Delete(id int) error {
	var permission model.Permission
	err := s.storage.FindByID(&permission, id)
	if err != nil {
		return err
	}
	err = s.storage.Delete(&permission)
	if err != nil {
		return err
	}
	return nil
}
