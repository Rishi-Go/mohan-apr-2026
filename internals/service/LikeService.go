package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
)

type LikeService interface {
	InsertLike(res dto.LikeRequest) (models.Like, error)
	GetLike(page int, limit int, offset int, userid uuid.UUID, blogid uuid.UUID) ([]models.Like, *dto.Pagination, error)
	SelectLike(id uuid.UUID) (models.Like, error)
	DeleteLike(id uuid.UUID, userid uuid.UUID) error
}

type likeService struct {
	Repo repository.LikeRepo
}

func InitLikeService(Repo repository.LikeRepo) LikeService {
	return &likeService{Repo}
}

func (likeService likeService) InsertLike(res dto.LikeRequest) (models.Like, error) {
	return likeService.Repo.InsertLike(res)
}

func (likeService likeService) GetLike(page int, limit int, offset int, userid uuid.UUID, blogid uuid.UUID) ([]models.Like, *dto.Pagination, error) {
	result, count, err := likeService.Repo.GetLike(page, limit, offset, userid, blogid)
	if err != nil {
		return []models.Like{}, nil, err
	}
	return result, &dto.Pagination{Page: page, Limit: limit, Total: int(count)}, err
}

func (likeService likeService) SelectLike(id uuid.UUID) (models.Like, error) {
	return likeService.Repo.SelectLike(id)
}

func (likeService likeService) DeleteLike(id uuid.UUID, userid uuid.UUID) error {
	UserID, err := likeService.Repo.GetLikeUserID(id)
	if err != nil {
		return err
	}

	if UserID != userid {
		return errors.New("Access Denied")
	}
	return likeService.Repo.DeleteLike(id)
}
