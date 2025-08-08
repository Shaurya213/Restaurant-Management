package di

import (
	"github.com/Shaurya213/Restaurant-Management/internal/auth"
	"github.com/Shaurya213/Restaurant-Management/internal/service"
)

type AppDeps struct {
	AdminService service.AdminService
	MenuService  service.MenuService
	OrderService service.OrderService
	AuthMid      *auth.AuthInterceptor
}
