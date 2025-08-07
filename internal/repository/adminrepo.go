package repository

import (
	"context"

	"github.com/Shaurya213/Restaurant-Management/internal/models"
	"gorm.io/gorm"
)

type AdminRepo interface {
	CreateAdmin(ctx context.Context, admin *models.Admin) error
	GetAdminByName(ctx context.Context, name string) (*models.Admin, error)
}

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) AdminRepo {
	return &adminRepo{db: db}
}

func (r *adminRepo) CreateAdmin(ctx context.Context, admin *models.Admin) error {
	return r.db.WithContext(ctx).Create(admin).Error
}

func (r *adminRepo) GetAdminByName(ctx context.Context, name string) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
