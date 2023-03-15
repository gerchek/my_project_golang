package storage

import (
	"my_project/internal/model"

	"gorm.io/gorm"
)

type PermissionStorage interface {
	All() []*model.Permission
	Create(permission *model.Permission) (data *model.Permission, err error)
	Update(permission *model.Permission) (data *model.Permission, err error)
	FindByID(permission *model.Permission, id int) error
	Delete(permission *model.Permission) error
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

func (ps *permissionStorage) Create(permission *model.Permission) (data *model.Permission, err error) {
	if err := ps.client.Create(permission).Error; err != nil {
		return permission, err
	}
	return permission, nil
}

func (ps *permissionStorage) Update(permission *model.Permission) (data *model.Permission, err error) {
	if err := ps.client.Save(permission).Error; err != nil {
		return permission, err
	}
	return permission, nil
}

func (ps *permissionStorage) FindByID(permission *model.Permission, id int) error {
	if err := ps.client.First(permission, id).Error; err != nil {
		return err
	}
	return nil
}

func (ps *permissionStorage) Delete(permission *model.Permission) error {
	err := ps.client.Delete(permission).Error
	if err != nil {
		return err
	}
	return nil
}
