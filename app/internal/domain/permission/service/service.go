package service

import (
	"my_project/internal/domain/permission/dto"
	"my_project/internal/domain/permission/storage"
	"my_project/internal/model"

	"github.com/go-redis/redis/v8"
)

type PermissionService interface {
	All() []*model.Permission
	Create(permissionDTO *dto.PermissionDTO) (data *model.Permission, err error)
	Update(permissionDTO *dto.PermissionDTO, id int) (data *model.Permission, err error)
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

func (s *permissionService) Create(permissionDTO *dto.PermissionDTO) (data *model.Permission, err error) {
	permission := &model.Permission{
		Name: permissionDTO.Name,
	}
	data, err = s.storage.Create(permission)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *permissionService) Update(permissionDTO *dto.PermissionDTO, id int) (data *model.Permission, err error) {
	var oldPermission model.Permission
	// var data model.Permission
	err = s.storage.FindByID(&oldPermission, id)
	if err != nil {
		return nil, err
	}
	oldPermission.Name = permissionDTO.Name
	data, err = s.storage.Update(&oldPermission)
	if err != nil {
		return nil, err
	}
	return data, nil
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
