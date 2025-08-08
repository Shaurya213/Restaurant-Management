package service

import (
	"context"
	"time"

	"github.com/Shaurya213/Restaurant-Management/internal/models"
	"github.com/Shaurya213/Restaurant-Management/internal/repository"
)

type OrderService interface {
	PlaceOrder(ctx context.Context, order *models.Order) error
	MarkPaid(ctx context.Context, id uint) error
	MarkServed(ctx context.Context, id uint) error
	ListOrders(ctx context.Context, paidOnly bool) ([]models.Order, error)
}

type orderService struct {
	repo repository.OrderRepo
}

func NewOrderService(repo repository.OrderRepo) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) PlaceOrder(ctx context.Context, order *models.Order) error {
	var total float64
	for _, item := range order.Items {
		total += float64(item.Quantity) * item.Price
	}
	order.Total = total
	order.CreatedAt = time.Now()
	return s.repo.CreateOrder(ctx, order)
}

func (s *orderService) MarkPaid(ctx context.Context, id uint) error {
	return s.repo.MarkPaid(ctx, id)
}

func (s *orderService) MarkServed(ctx context.Context, id uint) error {
	return s.repo.MarkServed(ctx, id)
}

func (s *orderService) ListOrders(ctx context.Context, paidOnly bool) ([]models.Order, error) {
	return s.repo.ListOrders(ctx, paidOnly)
}
