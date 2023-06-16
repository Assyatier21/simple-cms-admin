package middleware

import (
	"errors"
	"fmt"
	"time"

	"cms-admin/config"
	m "cms-admin/models"

	"github.com/golang-jwt/jwt/v4"
)

func getJWTSecretKey() string {
	return config.JWT_SECRET_KEY
}

func GenerateToken(registry m.User) string {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = registry.Name
	claims["phone"] = registry.Phone
	claims["role"] = registry.Role
	claims["expired_at"] = time.Now().Add(60 * time.Minute)

	t, _ := token.SignedString([]byte(getJWTSecretKey()))

	return t
}

func ParseTokenJWT(tokenString string) (m.UserClaimsResp, error) {
	var userClaims = m.UserClaimsResp{}
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(getJWTSecretKey()), nil
	})
	if err != nil {
		return userClaims, err
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return userClaims, errors.New("invalid token claims")
	}

	expiredAtStr, _ := claims["expired_at"].(string)
	expiredAt, err := time.Parse(time.RFC3339, expiredAtStr)
	if err != nil {
		return userClaims, fmt.Errorf("failed to parse expired_at value: %v", err)
	}
	userClaims = m.UserClaimsResp{
		Name:      claims["name"].(string),
		Phone:     claims["phone"].(string),
		Role:      claims["role"].(string),
		ExpiredAt: expiredAt,
	}

	return userClaims, nil
}

func IsTokenExpired(t time.Time) bool {
	now := time.Now()
	return t.Before(now)
}
