package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AuthRepo interface {
	InsertSignUP(res dto.SignUpRequest) error
}

type authRepo struct {
	Db *gorm.DB
}

func InitAuthRepo(Db *gorm.DB) AuthRepo {
	return &authRepo{Db}
}

func (auth authRepo) InsertSignUP(res dto.SignUpRequest) error {
	auth_id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	row := models.BlogUsers{
		ID:           auth_id,
		FirstName:    res.FirstName,
		LastName:     res.LastName,
		UserName:     res.UserName,
		PasswordHash: res.Password,
		Email:        res.Email,
	}

	result := auth.Db.Create(&row)
	err = result.Error
	if err != nil {
		return err
	}
	return nil
}
