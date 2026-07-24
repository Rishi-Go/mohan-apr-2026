package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
)

type AuthService interface {
	InsertUser(res dto.SignUpRequest) error
	GetUser(page int, limit int, offset int, username string, email string) ([]models.BlogUsers, *dto.Pagination, error)
	SelectUser(id uuid.UUID) (models.BlogUsers, error)
	UpdateUser(res dto.SignUpRequest, id uuid.UUID) error
	DeleteUser(id uuid.UUID) error
}

type authService struct {
	Repo repository.AuthRepo
}

func InitAuthService(Repo repository.AuthRepo) AuthService {
	return &authService{Repo}
}

func (auth authService) InsertUser(res dto.SignUpRequest) error {
	return auth.Repo.InsertUser(res)
}

func (auth authService) GetUser(page int, limit int, offset int, username string, email string) ([]models.BlogUsers, *dto.Pagination, error) {
	return auth.Repo.GetUser(page, limit, offset, username, email)
}

func (auth authService) SelectUser(id uuid.UUID) (models.BlogUsers, error) {
	return auth.Repo.SelectUser(id)
}

func (auth authService) UpdateUser(res dto.SignUpRequest, id uuid.UUID) error {
	return auth.Repo.UpdateUser(res, id)
}
func (auth authService)DeleteUser(id uuid.UUID) error{
	return auth.Repo.DeleteUser(id)
}