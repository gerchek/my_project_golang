package storage

import (
	"fmt"
	"my_project/internal/domain/admin/dto"
	"my_project/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdminStorage interface {
	All() []*model.Admin
	FindByUsername(admin *model.Admin, username string) error
	Create(admin *model.Admin) error
	FindByID(admin *model.Admin, id int) error
	Update(admin *model.Admin, adminDTO *dto.AdminDTO) error
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
	var admins []*model.Admin
	as.client.Preload("MyComments").Preload("MyProducts.Comments").Preload("Roles.Permissions").Preload("Roles").Find(&admins)
	return admins
}

func (as *adminStorage) FindByUsername(admin *model.Admin, username string) error {
	if err := as.client.Where("username = ?", username).Preload("MyComments").Preload("MyProducts.Comments").Preload("Roles.Permissions").Preload("Roles").First(admin).Error; err != nil {
		return err
	}
	return nil
}

func (as *adminStorage) Create(admin *model.Admin) error {
	if err := as.client.Omit("Roles.*").Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (as *adminStorage) FindByID(admin *model.Admin, id int) error {
	if err := as.client.Preload(clause.Associations).First(admin, id).Error; err != nil {
		return err
	}
	return nil
}

func (as *adminStorage) Update(admin *model.Admin, adminDTO *dto.AdminDTO) error {
	// if err := as.client.Save(admin).Error; err != nil {
	// 	return err
	// }

	if err := as.client.Save(admin).Association("Roles").Append(adminDTO.Roles_append); err != nil {
		fmt.Println(err.Error())
	}
	if err := as.client.Model(&admin).Association("Roles").Delete(adminDTO.Roles_delete); err != nil {
		fmt.Println(err.Error())
	}
	return nil

}
