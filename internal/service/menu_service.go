package service

import (
	"context"

	"github.com/Shaurya213/Restaurant-Management/internal/models"
	"github.com/Shaurya213/Restaurant-Management/internal/repository"
)

type MenuService interface {
	AddItem(ctx context.Context, item *models.MenuItem) error
	GetMenu(ctx context.Context) ([]models.MenuItem, error)
	UpdateItem(ctx context.Context, item *models.MenuItem) error
	DeleteItem(ctx context.Context, id uint) error
}

type menuService struct {
	repo repository.MenuRepo
}

func NewMenuService(repo repository.MenuRepo) MenuService {
	return &menuService{repo: repo}
}

func (s *menuService) AddItem(ctx context.Context, item *models.MenuItem) error {
	return s.repo.CreateItem(ctx, item)
}

func (s *menuService) GetMenu(ctx context.Context) ([]models.MenuItem, error) {
	return s.repo.ListItems(ctx)
}

func (s *menuService) UpdateItem(ctx context.Context, item *models.MenuItem) error {
	return s.repo.UpdateItem(ctx, item)
}

func (s *menuService) DeleteItem(ctx context.Context, id uint) error {
	return s.repo.DeleteItem(ctx, id)
}
