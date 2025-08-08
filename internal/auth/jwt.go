package auth

import (
	"time"

	"github.com/Shaurya213/Restaurant-Management/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	AdminID uint   `json:"admin_id"`
	Name    string `json:"name"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secret []byte
	ttl    time.Duration
}

func NewJWTManager(cfg *config.Config) *JWTManager {
	return &JWTManager{
		secret: []byte(cfg.JWTSecret),
		ttl:    cfg.JWTExpiry,
	}
}

func (m *JWTManager) Generate(adminID uint, name string) (string, error) {
	now := time.Now()
	claims := &Claims{
		AdminID: adminID,
		Name:    name,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.ttl)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

func (m *JWTManager) Verify(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return m.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
