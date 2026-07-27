package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
)

type BlogService interface {
	InsertBlog(res dto.BlogRequest) error
	GetBlog(page int, limit int, offset int, title string, categoryId uuid.UUID, authorId uuid.UUID) ([]models.Blog, *dto.Pagination, error)
	SelectBlog(id uuid.UUID) (models.Blog, error)
	UpdateBlog(res dto.BlogRequest, id uuid.UUID) error
	DeleteBlog(id uuid.UUID) error
}

type blogService struct {
	Repo repository.BlogRepo
}

func InitBlogService(Repo repository.BlogRepo) BlogService {
	return &blogService{Repo}
}

func (blogService blogService) InsertBlog(res dto.BlogRequest) error {
	return blogService.Repo.InsertBlog(res)
}

func (blogService blogService) GetBlog(page int, limit int, offset int, title string, categoryId uuid.UUID, authorId uuid.UUID) ([]models.Blog, *dto.Pagination, error) {
	return blogService.Repo.GetBlog(page,limit,offset,title,categoryId,authorId)
}
func (blogService blogService)  SelectBlog(id uuid.UUID) (models.Blog, error){
	return blogService.Repo.SelectBlog(id)
}
func (blogService blogService) UpdateBlog(res dto.BlogRequest, id uuid.UUID) error{
	return blogService.Repo.UpdateBlog(res,id)
}
func (blogService blogService) DeleteBlog(id uuid.UUID) error{
	return blogService.Repo.DeleteBlog(id)
}