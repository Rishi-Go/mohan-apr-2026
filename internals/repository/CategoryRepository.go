package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	InsertCategory(res dto.CategoryRequest) error
}

type categoryRepo struct {
	Db *gorm.DB
}

func InitRepo(Db *gorm.DB) CategoryRepo {
	return &categoryRepo{Db}
}

func (category categoryRepo) InsertCategory(res dto.CategoryRequest) error {
	category_id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	row := models.Category{
		ID:           category_id,
		CategoryName: res.CategoryName,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
		DeletedAt:    res.DeletedAt,
	}

	result := category.Db.Create(&row)
	err = result.Error
	if err != nil {
		return err
	}
	return nil
}
