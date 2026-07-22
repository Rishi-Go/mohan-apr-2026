package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
)

type AuthService interface {
	InsertSignUP(res dto.SignUpRequest) error
}

type authService struct {
	Repo repository.AuthRepo
}

func InitAuthService(Repo repository.AuthRepo) AuthService {
	return &authService{Repo}
}

func (auth authService) InsertSignUP(res dto.SignUpRequest) error {
	return auth.Repo.InsertSignUP(res)
}
