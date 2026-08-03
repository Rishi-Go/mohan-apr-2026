package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/middleware"
	"blog_post/internals/repository"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUpUser(res dto.SignUpRequest) error
	GetUser(page int, limit int, offset int, username string, email string) ([]models.BlogUsers, *dto.Pagination, error)
	SelectUser(id uuid.UUID) (models.BlogUsers, error)
	UpdateUser(res dto.SignUpRequest, id uuid.UUID, userid uuid.UUID, role string) error
	DeleteUser(id uuid.UUID, userid uuid.UUID, role string) error

	LogInUser(res dto.LogInRequest) (string, error)
}

type authService struct {
	Repo repository.AuthRepo
}

func InitAuthService(Repo repository.AuthRepo) AuthService {
	return &authService{Repo}
}

func (auth authService) SignUpUser(res dto.SignUpRequest) error {
	return auth.Repo.SignUpUser(res)
}

func (auth authService) GetUser(page int, limit int, offset int, username string, email string) ([]models.BlogUsers, *dto.Pagination, error) {
	return auth.Repo.GetUser(page, limit, offset, username, email)
}

func (auth authService) SelectUser(id uuid.UUID) (models.BlogUsers, error) {
	return auth.Repo.SelectUser(id)
}

func (auth authService) UpdateUser(res dto.SignUpRequest, id uuid.UUID, userid uuid.UUID, role string) error {
	UserID, err := auth.Repo.GetBlogUserID(id)
	if err != nil {
		return err
	}

	if UserID != userid && role != "Admin" {
		return errors.New("Access Denied")
	}
	return auth.Repo.UpdateUser(res, id, userid, role)
}
func (auth authService) DeleteUser(id uuid.UUID, userid uuid.UUID, role string) error {
	UserID, err := auth.Repo.GetBlogUserID(id)
	if err != nil {
		return err
	}

	if UserID != userid && role != "Admin" {
		return errors.New("Access Denied")
	}

	return auth.Repo.DeleteUser(id, userid, role)
}

func (auth authService) LogInUser(res dto.LogInRequest) (string, error) {

	users, err := auth.Repo.LogInUser(res)

	if users.UserName != res.UserName {
		return "Invalid Username, check your username", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.PasswordHash), []byte(res.Password))
	if err != nil {
		return "", errors.New("Invalid Password, check your password")
	}

	TokenStr, err := middleware.GenerateToken(res, users)
	if err != nil {
		return "", errors.New("Failed to generate token")
	}
	return TokenStr, nil
}
