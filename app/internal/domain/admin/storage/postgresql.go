package storage

import (
	"my_project/internal/model"

	"gorm.io/gorm"
)

type AdminStorage interface {
	All() []*model.Admin
	FindByUsername(admin *model.Admin, username string) error
	Create(admin *model.Admin) error
}

type adminStorage struct {
	client *gorm.DB
}

func NewAdminStorage(client *gorm.DB) AdminStorage {
	return &adminStorage{
		client: client,
	}
}

func (as *adminStorage) All() []*model.Admin {
	// a := "domain/admin/service/All()"
	var admins []*model.Admin
	// res := as.client.Find(&admins)
	as.client.Find(&admins)
	return admins
}

func (as *adminStorage) FindByUsername(admin *model.Admin, username string) error {
	if err := as.client.Where("username = ?", username).First(admin).Error; err != nil {
		return err
	}
	return nil
}

func (as *adminStorage) Create(admin *model.Admin) error {
	if err := as.client.Create(admin).Error; err != nil {
		return err
	}
	return nil
}
