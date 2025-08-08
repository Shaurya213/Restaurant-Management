package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/Shaurya213/Restaurant-Management/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ctxKey string

const (
	AdminIDKey   ctxKey = "admin_id"
	AdminNameKey ctxKey = "admin_name"
)

var publicMethods = map[string]bool{
	"/api.v1.AuthService/Register": true,
	"/api.v1.AuthService/Login":    true,
}

type AuthInterceptor struct {
	jwt *JWTManager
}

func NewAuthInterceptor(cfg *config.Config) *AuthInterceptor {
	return &AuthInterceptor{jwt: NewJWTManager(cfg)}
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, next grpc.UnaryHandler) (interface{}, error) {

		if publicMethods[info.FullMethod] {
			return next(ctx, req)
		}

		token, err := extractBearer(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "missing/invalid authorization header")
		}

		claims, err := a.jwt.Verify(token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid or expired token")
		}

		// Attach to context for handlers
		ctx = context.WithValue(ctx, AdminIDKey, claims.AdminID)
		ctx = context.WithValue(ctx, AdminNameKey, claims.Name)

		return next(ctx, req)
	}
}

func extractBearer(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata")
	}
	values := md.Get("authorization")
	if len(values) == 0 {
		values = md.Get("Authorization")
		if len(values) == 0 {
			return "", errors.New("no authorization header")
		}
	}
	parts := strings.SplitN(values[0], " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.New("bad auth scheme")
	}
	return parts[1], nil
}
