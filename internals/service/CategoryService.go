package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
)

type CategoryService interface {
	InsertCategory(res dto.CategoryRequest) error
}

type categoryService struct {
	Repo repository.CategoryRepo
}

func InitService(Repo repository.CategoryRepo) CategoryService {
	return &categoryService{Repo}
}

func (category categoryService) InsertCategory(res dto.CategoryRequest) error {
	return category.Repo.InsertCategory(res)
}
