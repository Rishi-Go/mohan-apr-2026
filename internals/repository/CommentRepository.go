package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type CommentRepo interface {
	InsertComment(res dto.CommentRequest) error
	GetComment(page int, limit int, offset int, comment string, userid uuid.UUID, blogid uuid.UUID) ([]models.Comment, *dto.Pagination, error)
	SelectComment(id uuid.UUID) (models.Comment, error)
	UpdateComment(res dto.CommentRequest, id uuid.UUID) error
	DeleteComment(id uuid.UUID) error
}

type commentRepo struct {
	Db *gorm.DB
}

func InitCommentRepo(Db *gorm.DB) CommentRepo {
	return &commentRepo{Db}
}

func (comment commentRepo) InsertComment(res dto.CommentRequest) error {

	CommentId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	row := models.Comment{
		ID:      CommentId,
		Comment: res.Comment,
		BlogID:  res.BlogID,
		UserID:  res.UserID,
	}

	result := comment.Db.Create(&row)
	err = result.Error
	if err != nil {
		return err
	}
	return nil
}

func (commentRepo commentRepo) GetComment(page int, limit int, offset int, comment string, userid uuid.UUID, blogid uuid.UUID) ([]models.Comment, *dto.Pagination, error) {

	var comments []models.Comment

	var count int64

	query := commentRepo.Db.Model(&comments)

	err := query.Count(&count).Error
	if err != nil {
		return nil, nil, err
	}

	if comment != "" {
		records := query.Where("comment LIKE ?", "%"+comment+"%").Session(&gorm.Session{})
		if records.Error != nil {
			return nil, nil, records.Error
		}
	}

	if userid != uuid.Nil {
		record := query.Where("user_id = ?", userid).Session(&gorm.Session{})
		if record.Error != nil {
			return nil, nil, record.Error
		}
	}

	if blogid != uuid.Nil {
		record := query.Where("blog_id = ?", blogid).Session(&gorm.Session{})
		if record.Error != nil {
			return nil, nil, record.Error
		}
	}

	result := query.Limit(limit).Offset(offset).Find(&comments)

	if result.RowsAffected == 0 {
		return nil, nil, errors.New("Comment Record data not found")
	}
	return comments, &dto.Pagination{Page: page, Limit: limit, Total: int(count), Offset: offset}, nil
}

func (commentRepo commentRepo) SelectComment(id uuid.UUID) (models.Comment, error) {

	var comments models.Comment

	result := commentRepo.Db.First(&comments, "id =?", id)
	if result.RowsAffected == 0 {
		return models.Comment{}, errors.New("Like Record data not found")
	}
	return comments, nil
}

func (commentRepo commentRepo) UpdateComment(res dto.CommentRequest, id uuid.UUID) error {

	var comments models.Comment

	result := commentRepo.Db.Model(&comments).Where("id = ?", id).Updates(models.Comment{
		Comment: res.Comment,
		BlogID: res.BlogID,
		UserID: res.UserID,
	})

	if result.RowsAffected == 0 {
		return errors.New("Like Record data not found")
	}
	return nil
}

func (commentRepo commentRepo) DeleteComment(id uuid.UUID) error {

	var comments models.Comment
	result := commentRepo.Db.Model(&comments).Delete(&comments, id)
	if result.RowsAffected == 0 {
		return errors.New("Like Record data not found")
	}
	return nil
}