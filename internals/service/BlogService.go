package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
	"blog_post/pkg/logger"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type BlogService interface {
	InsertBlog(res dto.BlogRequest) (models.Blog, error)
	GetBlog(page int, limit int, offset int, title string, categoryId uuid.UUID, authorId uuid.UUID) ([]models.Blog, *dto.Pagination, error)
	SelectBlog(id uuid.UUID) (models.Blog, error)
	UpdateBlog(res dto.BlogRequest, id uuid.UUID, role string) error
	DeleteBlog(id uuid.UUID, authorID uuid.UUID, role string) error
}

type blogService struct {
	Repo repository.BlogRepo
}

func InitBlogService(Repo repository.BlogRepo) BlogService {
	return &blogService{Repo}
}

func (blogService blogService) InsertBlog(res dto.BlogRequest) (models.Blog, error) {
	return blogService.Repo.InsertBlog(res)
}

func (blogService blogService) GetBlog(page int, limit int, offset int, title string, categoryId uuid.UUID, authorId uuid.UUID) ([]models.Blog, *dto.Pagination, error) {
	result, count, err := blogService.Repo.GetBlog(page, limit, offset, title, categoryId, authorId)
	if err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return []models.Blog{}, nil, err
	}
	return result, &dto.Pagination{Page: page, Limit: limit, Total: int(count)}, err
}

func (blogService blogService) SelectBlog(id uuid.UUID) (models.Blog, error) {
	return blogService.Repo.SelectBlog(id)
}

func (blogService blogService) UpdateBlog(res dto.BlogRequest, id uuid.UUID, role string) error {

	AuthorID, err := blogService.Repo.GetBlogAuthorID(id)
	if err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return err
	}

	if AuthorID != res.AuthorID && role != "Admin" {
		logger.Log.Error("Access Denied")
		return errors.New("Access Denied")
	}
	return blogService.Repo.UpdateBlog(res, id, role)
}

func (blogService blogService) DeleteBlog(id uuid.UUID, authorID uuid.UUID, role string) error {

	AuthorID, err := blogService.Repo.GetBlogAuthorID(id)
	if err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return err
	}

	if AuthorID != authorID && role != "Admin" {
		logger.Log.Error("Access Denied")
		return errors.New("Access Denied")
	}

	return blogService.Repo.DeleteBlog(id, authorID, role)
}
