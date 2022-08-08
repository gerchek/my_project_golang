package service

import (
	"errors"
	"my_project/internal/domain/role/dto"
	"my_project/internal/domain/role/storage"
	"my_project/internal/model"

	"github.com/go-redis/redis/v8"
)

type RoleService interface {
	All() []*model.Role
	Create(roleDTO *dto.RoleDTO) error
	Update(roleDTO *dto.RoleDTO, id int) error
	Delete(id int) error
}

type roleService struct {
	storage     storage.RoleStorage
	redisClient *redis.Client
}

func NewRoleService(storage storage.RoleStorage, redisClient *redis.Client) RoleService {
	return &roleService{
		storage:     storage,
		redisClient: redisClient,
	}
}

func (s *roleService) All() []*model.Role {
	return s.storage.All()
}

func (s *roleService) Create(roleDTO *dto.RoleDTO) error {
	admin := &model.Role{
		Name:        roleDTO.Name,
		Permissions: roleDTO.Permissions_append,
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

func (s *roleService) Update(roleDTO *dto.RoleDTO, id int) error {
	var oldRole model.Role
	err := s.storage.FindByID(&oldRole, id)
	if err != nil {
		return err
	}
	oldRole.Name = roleDTO.Name
	err = s.storage.Update(&oldRole, roleDTO)
	if err != nil {
		return err
	}
	return nil
}

func (s *roleService) Delete(id int) error {
	var role model.Role
	err := s.storage.FindByID(&role, id)
	if err != nil {
		return err
	}
	err = s.storage.Delete(&role)
	if err != nil {
		return err
	}
	return nil
}
