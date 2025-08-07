package repository

import (
	"context"

	"github.com/Shaurya213/Restaurant-Management/internal/models"
	"gorm.io/gorm"
)

type MenuRepo interface {
	CreateItem(ctx context.Context, item *models.MenuItem) error
	ListItems(ctx context.Context) ([]models.MenuItem, error)
	UpdateItem(ctx context.Context, item *models.MenuItem) error
	DeleteItem(ctx context.Context, id uint) error
}

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) MenuRepo {
	return &menuRepo{db: db}
}

func (r *menuRepo) CreateItem(ctx context.Context, item *models.MenuItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *menuRepo) ListItems(ctx context.Context) ([]models.MenuItem, error) {
	var items []models.MenuItem
	err := r.db.WithContext(ctx).Find(&items).Error
	return items, err
}

func (r *menuRepo) UpdateItem(ctx context.Context, item *models.MenuItem) error {
	return r.db.WithContext(ctx).Save(item).Error
}

func (r *menuRepo) DeleteItem(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.MenuItem{}, id).Error
}
