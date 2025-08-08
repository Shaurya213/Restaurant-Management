//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/Shaurya213/Restaurant-Management/internal/auth"
	"github.com/Shaurya213/Restaurant-Management/internal/config"
	"github.com/Shaurya213/Restaurant-Management/internal/db"
	"github.com/Shaurya213/Restaurant-Management/internal/repository"
	"github.com/Shaurya213/Restaurant-Management/internal/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeApp() (*AppDeps, error) {
	gormDB := db.ProvideDB()
	adminRepo := repository.NewAdminRepo(gormDB)
	adminService := service.NewAdminService(adminRepo)
	menuRepo := repository.NewMenuRepo(gormDB)
	menuService := service.NewMenuService(menuRepo)
	orderRepo := repository.NewOrderRepo(gormDB)
	orderService := service.NewOrderService(orderRepo)
	configConfig := config.Load()
	authInterceptor := auth.NewAuthInterceptor(configConfig)
	appDeps := &AppDeps{
		AdminService: adminService,
		MenuService:  menuService,
		OrderService: orderService,
		AuthMid:      authInterceptor,
	}
	return appDeps, nil
}

// wire.go:

// Provider sets
var configSet = wire.NewSet(config.Load)

var dbSet = wire.NewSet(db.ProvideDB)

var repoSet = wire.NewSet(repository.NewAdminRepo, repository.NewMenuRepo, repository.NewOrderRepo)

var serviceSet = wire.NewSet(service.NewAdminService, service.NewMenuService, service.NewOrderService)

var authSet = wire.NewSet(auth.NewAuthInterceptor)
