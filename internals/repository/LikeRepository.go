package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type LikeRepo interface {
	InsertLike(res dto.LikeRequest) error
	// GetLike(page int, limit int, offset int, like bool,user_id uuid.UUID) ([]models.Like, *dto.Pagination, error)
}

type likeRepo struct {
	Db *gorm.DB
}

func InitLikeRepo(Db *gorm.DB) LikeRepo {
	return &likeRepo{Db}
}

func (like likeRepo) InsertLike(res dto.LikeRequest) error {

	LikeId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	row := models.Like{
		ID:           LikeId,
		LikeResponse: res.LikeResponse,
		UserID:       res.UserID,
	}

	result := like.Db.Create(&row)
	err = result.Error
	if err != nil {
		return err
	}
	return nil
}

// func (like likeRepo) GetLike(page int, limit int, offset int, like bool,user_id uuid.UUID) ([]models.Like, *dto.Pagination, error) {

// 	var likes []models.Like

// 	var count int64

// 	query := like.Db.Model(&likes)

// 	err := query.Count(&count).Error
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	if like != false {
// 		records := query.Where(" ").Session(&gorm.Session{})
// 		if records.Error != nil {
// 			return nil, nil, records.Error
// 		}
// 	}

// 	result := query.Limit(limit).Offset(offset).Find(&likes)

// 	if result.RowsAffected == 0 {
// 		return nil, nil, errors.New("Like Record data not found")
// 	}
// 	return likes, &dto.Pagination{Page: page, Limit: limit, Total: int(count), Offset: offset}, nil
// }

// func (repo likeRepo) SelectLike(id uuid.UUID) (models.Like, error) {

// 	var likes models.Like

// 	result := repo.Db.First(&likes, "id =?", id)
// 	if result.RowsAffected == 0 {
// 		return models.Like{}, errors.New("Like Record data not found")
// 	}
// 	return likes, nil
// }
