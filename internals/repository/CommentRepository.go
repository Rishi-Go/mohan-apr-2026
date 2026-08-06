package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type CommentRepo interface {
	InsertComment(res dto.CommentRequest) (models.Comment, error)
	GetComment(page int, limit int, offset int, comment string, userid uuid.UUID, blogid uuid.UUID) ([]models.Comment, int, error)
	SelectComment(id uuid.UUID) (models.Comment, error)
	UpdateComment(res dto.CommentRequest, id uuid.UUID) error
	DeleteComment(id uuid.UUID, userid uuid.UUID) error

	GetCommentUserID(id uuid.UUID) (uuid.UUID, error)
}

type commentRepo struct {
	Db *gorm.DB
}

func InitCommentRepo(Db *gorm.DB) CommentRepo {
	return &commentRepo{Db}
}

func (comment commentRepo) InsertComment(res dto.CommentRequest) (models.Comment, error) {

	CommentId, err := uuid.NewV7()
	if err != nil {
		return models.Comment{}, err
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
		return models.Comment{}, err
	}
	return row, nil
}

func (commentRepo commentRepo) GetComment(page int, limit int, offset int, comment string, userid uuid.UUID, blogid uuid.UUID) ([]models.Comment, int, error) {

	var comments []models.Comment

	var count int64

	query := commentRepo.Db.Model(&comments)

	err := query.Count(&count).Error
	if err != nil {
		return []models.Comment{}, 0, err
	}

	if comment != "" {
		records := query.Where("comment ILIKE ?", "%"+comment+"%").Session(&gorm.Session{})
		if records.Error != nil {
			return []models.Comment{}, 0, records.Error
		}
	}

	if userid != uuid.Nil {
		record := query.Where("user_id = ?", userid).Session(&gorm.Session{})
		if record.Error != nil {
			return []models.Comment{}, 0, record.Error
		}
	}

	if blogid != uuid.Nil {
		record := query.Where("blog_id = ?", blogid).Session(&gorm.Session{})
		if record.Error != nil {
			return []models.Comment{}, 0, record.Error
		}
	}

	result := query.Limit(limit).Offset(offset).Find(&comments)

	if result.RowsAffected == 0 {
		return []models.Comment{}, 0, errors.New("Comment Record data not found")
	}
	return comments, int(count), nil
}

func (commentRepo commentRepo) SelectComment(id uuid.UUID) (models.Comment, error) {

	var comments models.Comment

	result := commentRepo.Db.First(&comments, "id =?", id)
	if result.RowsAffected == 0 {
		return models.Comment{}, errors.New("Comment Record data not found")
	}
	return comments, nil
}

func (commentRepo commentRepo) UpdateComment(res dto.CommentRequest, id uuid.UUID) error {

	var comments models.Comment

	result := commentRepo.Db.Model(&comments).Where("id = ?", id).Updates(models.Comment{
		Comment: res.Comment,
		BlogID:  res.BlogID,
		UserID:  res.UserID,
	})

	if result.RowsAffected == 0 {
		return errors.New("Comment Record data not found")
	}
	return nil
}

func (commentRepo commentRepo) DeleteComment(id uuid.UUID, userid uuid.UUID) error {

	var comments models.Comment
	result := commentRepo.Db.Model(&comments).Delete(&comments, id)
	if result.RowsAffected == 0 {
		return errors.New("Comment Record data not found")
	}
	return nil
}

func (commentRepo commentRepo) GetCommentUserID(id uuid.UUID) (uuid.UUID, error) {

	var Comments models.Comment

	result := commentRepo.Db.First(&Comments, "id =?", id)
	if result.RowsAffected == 0 {
		return uuid.Nil, errors.New("Comments Record data not found")
	}

	UserId := Comments.UserID

	return UserId, nil
}
