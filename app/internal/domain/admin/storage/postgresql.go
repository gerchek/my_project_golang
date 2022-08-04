package storage

import (
	"gorm.io/gorm"
)

type AdminStorage interface {
	All() string
}

type adminStorage struct {
	client *gorm.DB
}

func NewAdminStorage(client *gorm.DB) AdminStorage {
	return &adminStorage{
		client: client,
	}
}

func (as *adminStorage) All() string {
	a := "domain/admin/service/All()"
	return a
}
