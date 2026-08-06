package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type LikeRepo interface {
	InsertLike(res dto.LikeRequest) (models.Like, error)
	GetLike(page int, limit int, offset int, user_id uuid.UUID, blogid uuid.UUID) ([]models.Like, int, error)
	SelectLike(id uuid.UUID) (models.Like, error)
	DeleteLike(id uuid.UUID) error

	GetLikeUserID(id uuid.UUID) (uuid.UUID, error)
}

type likeRepo struct {
	Db *gorm.DB
}

func InitLikeRepo(Db *gorm.DB) LikeRepo {
	return &likeRepo{Db}
}

func (likeRepo likeRepo) InsertLike(res dto.LikeRequest) (models.Like, error) {

	LikeId, err := uuid.NewV7()
	if err != nil {
		return models.Like{}, err
	}

	row := models.Like{
		ID:           LikeId,
		LikeResponse: res.LikeResponse,
		UserID:       res.UserID,
		BlogID:       res.BlogID,
	}

	result := likeRepo.Db.Create(&row)
	err = result.Error
	if err != nil {
		return models.Like{}, err
	}
	return row, nil
}

func (likeRepo likeRepo) GetLike(page int, limit int, offset int, userid uuid.UUID, blogid uuid.UUID) ([]models.Like, int, error) {

	var likes []models.Like

	var count int64

	query := likeRepo.Db.Model(&likes)

	err := query.Count(&count).Error
	if err != nil {
		return []models.Like{}, 0, err
	}

	if userid != uuid.Nil {
		record := query.Where("user_id = ?", userid).Session(&gorm.Session{})
		if record.RowsAffected == 0 {
			return []models.Like{}, 0, errors.New("Invalid userid or userid not found")
		}
	}

	if blogid != uuid.Nil {
		record := query.Where("blog_id = ?", blogid).Session(&gorm.Session{})
		if record.RowsAffected == 0 {
			return []models.Like{}, 0, errors.New("Invalid blogid or blogid not found")
		}
	}
	result := query.Limit(limit).Offset(offset).Find(&likes)

	if result.RowsAffected == 0 {
		return []models.Like{}, 0, errors.New("Like Record data not found")
	}
	return likes, int(count), nil
}

func (likeRepo likeRepo) SelectLike(id uuid.UUID) (models.Like, error) {

	var likes models.Like

	result := likeRepo.Db.First(&likes, "id =?", id)
	if result.RowsAffected == 0 {
		return models.Like{}, errors.New("Like Record data not found")
	}
	return likes, nil
}

func (likeRepo likeRepo) DeleteLike(id uuid.UUID) error {

	var likes models.Like
	result := likeRepo.Db.Model(&likes).Delete(&likes, id)
	if result.RowsAffected == 0 {
		return errors.New("Like Record data not found")
	}
	return nil
}

func (likeRepo likeRepo) GetLikeUserID(id uuid.UUID) (uuid.UUID, error) {

	var likes models.Like

	result := likeRepo.Db.First(&likes, "id =?", id)
	if result.RowsAffected == 0 {
		return uuid.Nil, errors.New("Reply Record data not found")
	}

	UserId := likes.UserID

	return UserId, nil
}
