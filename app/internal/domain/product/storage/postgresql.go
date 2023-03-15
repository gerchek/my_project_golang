package storage

import (
	"my_project/internal/domain/product/dto"
	"my_project/internal/model"

	"gorm.io/gorm"
)

type ProductStorage interface {
	All() []*model.Product
	Create(product *model.Product) error
	Update(product *model.Product, productDTO *dto.ProductDTO) error
	FindByID(product *model.Product, id int) error
	Delete(product *model.Product) error
}

type productStorage struct {
	client *gorm.DB
}

func NewProductStorage(client *gorm.DB) ProductStorage {
	return &productStorage{
		client: client,
	}
}

func (ps *productStorage) All() []*model.Product {
	var products []*model.Product
	ps.client.Preload("Comments").Find(&products)
	return products
}

func (ps *productStorage) Create(product *model.Product) error {
	if err := ps.client.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (ps *productStorage) FindByID(product *model.Product, id int) error {
	if err := ps.client.First(product, id).Error; err != nil {
		return err
	}
	return nil
}

func (ps *productStorage) Update(product *model.Product, productDTO *dto.ProductDTO) error {
	if err := ps.client.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (ps *productStorage) Delete(product *model.Product) error {
	err := ps.client.Delete(product).Error
	if err != nil {
		return err
	}
	return nil
}
