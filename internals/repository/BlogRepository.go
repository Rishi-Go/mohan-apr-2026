package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type BlogRepo interface {
	InsertBlog(res dto.BlogRequest) (models.Blog, error)
	GetBlog(page int, limit int, offset int, title string, categoryId uuid.UUID, authorId uuid.UUID) ([]models.Blog, int, error)
	SelectBlog(id uuid.UUID) (models.Blog, error)
	UpdateBlog(res dto.BlogRequest, id uuid.UUID, role string) error
	DeleteBlog(id uuid.UUID, authorID uuid.UUID, role string) error
	GetBlogAuthorID(id uuid.UUID) (uuid.UUID, error)
}

type blogRepo struct {
	Db *gorm.DB
}

func InitBlogRepo(Db *gorm.DB) BlogRepo {
	return &blogRepo{Db}
}

func (blogRepo blogRepo) InsertBlog(res dto.BlogRequest) (models.Blog, error) {

	BlogId, err := uuid.NewV7()
	if err != nil {
		return models.Blog{}, err
	}

	row := models.Blog{
		ID:         BlogId,
		Title:      res.Title,
		Content:    res.Content,
		CategoryId: res.CategoryId,
		AuthorID:   res.AuthorID,
	}
	result := blogRepo.Db.Create(&row)
	err = result.Error
	if err != nil {
		return models.Blog{}, err
	}
	return row, nil
}

func (blogRepo blogRepo) GetBlog(page int, limit int, offset int, title string, categoryId uuid.UUID, authorId uuid.UUID) ([]models.Blog, int, error) {

	var blogs []models.Blog

	var count int64

	query := blogRepo.Db.Model(&blogs)

	err := query.Count(&count).Error
	if err != nil {
		return []models.Blog{}, 0, err
	}

	if title != "" {
		records := query.Where("title ILIKE ?", "%"+title+"%").Session(&gorm.Session{})
		if records.Error != nil {
			return []models.Blog{}, 0, records.Error
		}
	}

	if categoryId != uuid.Nil {
		records := query.Where("category_id = ?", categoryId).Session(&gorm.Session{})
		if records.Error != nil {
			return []models.Blog{}, 0, records.Error
		}
	}

	if authorId != uuid.Nil {
		record := query.Where("author_id = ?", authorId).Session(&gorm.Session{})
		if record.Error != nil {
			return []models.Blog{}, 0, record.Error
		}
	}

	result := query.Limit(limit).Offset(offset).Find(&blogs)

	if result.RowsAffected == 0 {
		return []models.Blog{}, 0, errors.New("Blog Record data not found")
	}
	return blogs, int(count), nil
}

func (blogRepo blogRepo) SelectBlog(id uuid.UUID) (models.Blog, error) {

	var Blogs models.Blog

	result := blogRepo.Db.First(&Blogs, "id =?", id)
	if result.RowsAffected == 0 {
		return models.Blog{}, errors.New("Blog Record data not found")
	}
	return Blogs, nil
}

func (blogRepo blogRepo) UpdateBlog(res dto.BlogRequest, id uuid.UUID, role string) error {

	var blogs models.Blog

	result := blogRepo.Db.Model(&blogs).Where("id = ?", id).Updates(models.Blog{
		Title:      res.Title,
		Content:    res.Content,
		CategoryId: res.CategoryId,
		AuthorID:   res.AuthorID,
	})

	if result.RowsAffected == 0 {
		return errors.New("Blog Record data not found")
	}
	return nil
}

func (blogRepo blogRepo) DeleteBlog(id uuid.UUID, authorID uuid.UUID, role string) error {

	var blogs models.Blog
	result := blogRepo.Db.Model(&blogs).Delete(&blogs, id)
	if result.RowsAffected == 0 {
		return errors.New("Blog Record data not found")
	}
	return nil
}

func (blogRepo blogRepo) GetBlogAuthorID(id uuid.UUID) (uuid.UUID, error) {

	var Blogs models.Blog

	result := blogRepo.Db.First(&Blogs, "id =?", id)
	if result.RowsAffected == 0 {
		return uuid.Nil, errors.New("Blog Record data not found")
	}

	AuthorId := Blogs.AuthorID

	return AuthorId, nil
}
