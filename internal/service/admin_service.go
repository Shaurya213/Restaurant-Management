package service

import (
	"context"
	"errors"

	"github.com/Shaurya213/Restaurant-Management/internal/models"
	"github.com/Shaurya213/Restaurant-Management/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	Register(ctx context.Context, name, password string) error
	Login(ctx context.Context, name, password string) (*models.Admin, error)
}

type adminService struct {
	repo repository.AdminRepo
}

func NewAdminService(repo repository.AdminRepo) AdminService {
	return &adminService{repo: repo}
}

func (s *adminService) Register(ctx context.Context, name, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin := &models.Admin{
		Name:     name,
		Password: string(hashed),
	}
	return s.repo.CreateAdmin(ctx, admin)
}

func (s *adminService) Login(ctx context.Context, name, password string) (*models.Admin, error) {
	admin, err := s.repo.GetAdminByName(ctx, name)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return admin, nil
}
