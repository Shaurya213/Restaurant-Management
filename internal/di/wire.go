//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Shaurya213/Restaurant-Management/internal/auth"
	"github.com/Shaurya213/Restaurant-Management/internal/config"
	"github.com/Shaurya213/Restaurant-Management/internal/db"
	"github.com/Shaurya213/Restaurant-Management/internal/repository"
	"github.com/Shaurya213/Restaurant-Management/internal/service"
	"github.com/google/wire"
)

// Provider sets
var configSet = wire.NewSet(
	config.Load, // *config.Config
)

var dbSet = wire.NewSet(
	db.ProvideDB, // *gorm.DB (after InitMySQL)
)

var repoSet = wire.NewSet(
	repository.NewAdminRepo,
	repository.NewMenuRepo,
	repository.NewOrderRepo,
)

var serviceSet = wire.NewSet(
	service.NewAdminService,
	service.NewMenuService,
	service.NewOrderService,
)

var authSet = wire.NewSet(
	auth.NewAuthInterceptor,
)

func InitializeApp() (*AppDeps, error) {
	wire.Build(
		configSet,
		dbSet,
		repoSet,
		serviceSet,
		authSet,
		wire.Struct(new(AppDeps), "AdminService", "MenuService", "OrderService", "AuthMid"),
	)
	return nil, nil
}
