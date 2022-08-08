package service

import (
	"errors"
	"my_project/internal/domain/product/dto"
	"my_project/internal/domain/product/storage"
	"my_project/internal/model"

	"github.com/go-redis/redis/v8"
)

type ProductService interface {
	All() []*model.Product
	Create(productDTO *dto.ProductDTO) error
	Update(productDTO *dto.ProductDTO, id int) error
	Delete(id int) error
}

type productService struct {
	storage     storage.ProductStorage
	redisClient *redis.Client
}

func NewProductService(storage storage.ProductStorage, redisClient *redis.Client) ProductService {
	return &productService{
		storage:     storage,
		redisClient: redisClient,
	}
}

func (s *productService) All() []*model.Product {
	// a := s.storage.Test()
	return s.storage.All()
}

func (s *productService) Create(productDTO *dto.ProductDTO) error {
	product := &model.Product{
		AdminID: int(productDTO.AdminID),
		Name:    productDTO.Name,
	}

	err := s.storage.Create(product)
	if err != nil {
		return err
	}
	if err != nil {
		return errors.New("admin created but there was problem when updating order number")
	}
	return nil
}

func (s *productService) Update(productDTO *dto.ProductDTO, id int) error {
	var oldProduct model.Product
	err := s.storage.FindByID(&oldProduct, id)
	if err != nil {
		return err
	}
	oldProduct.Name = productDTO.Name
	err = s.storage.Update(&oldProduct, productDTO)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) Delete(id int) error {
	var product model.Product
	err := s.storage.FindByID(&product, id)
	if err != nil {
		return err
	}
	err = s.storage.Delete(&product)
	if err != nil {
		return err
	}
	return nil
}
