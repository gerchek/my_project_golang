package service

import (
	"context"
	"errors"
	"my_project/internal/domain/admin/dto"
	"my_project/internal/domain/admin/storage"
	"my_project/internal/model"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	All() []*model.Admin
	FindByUsername(username string) (*model.Admin, error)
	Create(adminDTO *dto.AdminDTO) error

	CreateAuth(userID string, td *model.TokenDetails) error
	DeleteAuth(accessUuid string) (int64, error)
}

type adminService struct {
	storage     storage.AdminStorage
	redisClient *redis.Client
}

func NewAdminService(storage storage.AdminStorage, redisClient *redis.Client) AdminService {
	return &adminService{
		storage:     storage,
		redisClient: redisClient,
	}
}

func (s *adminService) All() []*model.Admin {
	return s.storage.All()
}

func (s *adminService) FindByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	err := s.storage.FindByUsername(&admin, username)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

//Admin user CREATE
func (s *adminService) Create(adminDTO *dto.AdminDTO) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(*adminDTO.Password), 10)
	admin := &model.Admin{
		Username:  adminDTO.Username,
		FirstName: adminDTO.FirstName,
		LastName:  adminDTO.LastName,
		Password:  string(password),
		// Roles:       adminDTO.Roles,
		// Permissions: adminDTO.Permissions,
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

func (s *adminService) CreateAuth(userID string, td *model.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := s.redisClient.Set(context.Background(), td.AccessUuid, userID, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := s.redisClient.Set(context.Background(), td.RefreshUuid, userID, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil

}

func (s *adminService) DeleteAuth(accessUuid string) (int64, error) {
	deleted, _ := s.redisClient.Del(context.Background(), accessUuid).Result()
	if deleted == 0 {
		return 0, errors.New("key doesn't exists")
	}
	return deleted, nil
}
