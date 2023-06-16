package usecase

import (
	"cms-admin/middleware"
	m "cms-admin/models"
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Login(ctx context.Context, phone string, password string) (interface{}, error) {
	var (
		userJWT      = m.UserJWT{}
		userRegistry = m.User{}
		token        string
		err          error
	)

	userRegistry, err = u.repository.GetUserRegistry(ctx, phone, password)
	if err != nil {
		log.Println("[Usecase][User][LoginUser] failed to get user information, err: ", err)
		return userRegistry, nil
	}

	if userRegistry.Phone == "" {
		return userRegistry, errors.New("phone or password is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRegistry.Password), []byte(password))
	if err != nil {
		return userRegistry, errors.New("phone or password is incorrect")
	}

	token = middleware.GenerateToken(userRegistry)
	userJWT = m.UserJWT{
		Phone: userRegistry.Phone,
		Token: token,
	}

	return userJWT, nil
}
