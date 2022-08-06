package storage

import (
	"my_project/internal/model"

	"gorm.io/gorm"
)

type PermissionStorage interface {
	All() []*model.Permission
	Create(permission *model.Permission) error
}

type permissionStorage struct {
	client *gorm.DB
}

func NewPermissionStorage(client *gorm.DB) PermissionStorage {
	return &permissionStorage{
		client: client,
	}
}

func (ps *permissionStorage) All() []*model.Permission {
	permissions := make([]*model.Permission, 0)
	ps.client.Preload("Roles").Find(&permissions)
	return permissions
}

func (ps *permissionStorage) Create(permission *model.Permission) error {
	if err := ps.client.Create(permission).Error; err != nil {
		return err
	}
	return nil
}
