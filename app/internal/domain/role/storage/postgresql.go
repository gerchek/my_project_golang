package storage

import (
	"my_project/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleStorage interface {
	All() []*model.Role
	Create(role *model.Role) error
	Update(role *model.Role) error
	Delete(role *model.Role) error
	FindByID(role *model.Role, id int) error
}

type roleStorage struct {
	client *gorm.DB
}

func NewRoleStorage(client *gorm.DB) RoleStorage {
	return &roleStorage{
		client: client,
	}
}

func (rs *roleStorage) All() []*model.Role {
	var roles []*model.Role
	rs.client.Preload("Admins").Find(&roles)
	return roles
}

func (rs *roleStorage) Create(role *model.Role) error {
	if err := rs.client.Create(role).Error; err != nil {
		return err
	}
	return nil
}

func (rs *roleStorage) FindByID(role *model.Role, id int) error {
	if err := rs.client.First(role, id).Error; err != nil {
		return err
	}
	return nil
}

func (rs *roleStorage) Update(role *model.Role) error {
	if err := rs.client.Save(role).Error; err != nil {
		return err
	}
	return nil
}

func (rs *roleStorage) Delete(role *model.Role) error {
	err := rs.client.Select(clause.Associations).Delete(role).Error
	if err != nil {
		return err
	}
	return nil
}
