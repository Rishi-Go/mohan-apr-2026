package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
)

type CategoryService interface {
	InsertCategory(res dto.CategoryRequest) (models.Category, error)
	GetCategory(page int, limit int, offset int, name string) ([]models.Category, *dto.Pagination, error)
	SelectCategory(id uuid.UUID) (models.Category, error)
	UpdateCategory(res dto.CategoryRequest, id uuid.UUID) error
	DeleteCategory(id uuid.UUID) error
}

type categoryService struct {
	Repo repository.CategoryRepo
}

func InitCategoryService(Repo repository.CategoryRepo) CategoryService {
	return &categoryService{Repo}
}

func (category categoryService) InsertCategory(res dto.CategoryRequest) (models.Category, error) {
	return category.Repo.InsertCategory(res)
}

func (category categoryService) GetCategory(page int, limit int, offset int, name string) ([]models.Category, *dto.Pagination, error) {
	return category.Repo.GetCategory(page, limit, offset, name)
}

func (category categoryService) SelectCategory(id uuid.UUID) (models.Category, error) {
	return  category.Repo.SelectCategory(id)
}

func (category categoryService) UpdateCategory(res dto.CategoryRequest, id uuid.UUID) error {
	return  category.Repo.UpdateCategory(res,id)
}

func (category categoryService) DeleteCategory(id uuid.UUID) error{
	return category.Repo.DeleteCategory(id)
}
