package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	User struct {
		Phone    string `json:"phone"`
		Name     string `json:"name"`
		Role     string `json:"role"`
		Password string `json:"password"`
	}

	GetUserReq struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	UserJWT struct {
		Phone string `json:"phone"`
		Token string `json:"token"`
	}

	UserClaims struct {
		Name      string               `json:"name"`
		Phone     string               `json:"phone"`
		Role      string               `json:"role"`
		ExpiredAt time.Time            `json:"expired_at"`
		Claims    jwt.RegisteredClaims `json:"claims"`
	}
	UserClaimsResp struct {
		Name      string    `json:"name"`
		Phone     string    `json:"phone"`
		Role      string    `json:"role"`
		ExpiredAt time.Time `json:"expired_at"`
	}
)
