package repository

import (
	"context"

	"github.com/Shaurya213/Restaurant-Management/internal/models"
	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, order *models.Order) error
	MarkPaid(ctx context.Context, id uint) error
	MarkServed(ctx context.Context, id uint) error
	ListOrders(ctx context.Context, onlyPaid bool) ([]models.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db: db}
}

func (r *orderRepo) CreateOrder(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepo) MarkPaid(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", id).Update("is_paid", true).Error
}

func (r *orderRepo) MarkServed(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", id).Update("is_served", true).Error
}

func (r *orderRepo) ListOrders(ctx context.Context, onlyPaid bool) ([]models.Order, error) {
	var orders []models.Order
	query := r.db.WithContext(ctx).Preload("Items")
	if onlyPaid {
		query = query.Where("is_paid = ?", true)
	}
	err := query.Find(&orders).Error
	return orders, err
}
