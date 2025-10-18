package utils

import (
	"errors"
	"time"

	"github.com/alwijein/be_football/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims represents JWT claims
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken generates JWT token
func GenerateToken(userID uint, username string) (string, error) {
	cfg := config.GlobalConfig

	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(cfg.JWT.ExpirationHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// ValidateToken validates JWT token
func ValidateToken(tokenString string) (*JWTClaims, error) {
	cfg := config.GlobalConfig

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
