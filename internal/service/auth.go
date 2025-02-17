package service

import (
	"errors"
	"github.com/mdvasilyev/avito-shop/internal/helper"
	"github.com/mdvasilyev/avito-shop/internal/repository"
	"log/slog"
)

type AuthService struct {
	lgr  *slog.Logger
	repo *repository.AuthRepository
}

func NewAuthService(lgr *slog.Logger, repo *repository.AuthRepository) *AuthService {
	return &AuthService{lgr: lgr, repo: repo}
}

func (srv *AuthService) Auth(username, password string) (string, error) {
	srv.lgr.Info("Authing")

	var userID int
	user, err := srv.repo.GetUser(username)
	if err != nil {
		hashedPass, err := helper.HashPassword(password)
		if err != nil {
			srv.lgr.Error("Error hashing password")
			return "", errors.New("error hashing password")
		}

		userID, err = srv.repo.UpdateUser(username, hashedPass)
		if err != nil {
			srv.lgr.Error("Error updating user")
			return "", errors.New("error updating user")
		}
	} else if !helper.CheckHashedPassword(password, user.PassHash) {
		srv.lgr.Error("Wrong password")
		return "", errors.New("wrong password")
	}

	return helper.NewToken(userID)
}
