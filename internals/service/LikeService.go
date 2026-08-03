package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
)

type LikeService interface {
	InsertLike(res dto.LikeRequest) error
	GetLike(page int, limit int, offset int, like string, userid uuid.UUID, blogid uuid.UUID) ([]models.Like, *dto.Pagination, error)
	SelectLike(id uuid.UUID) (models.Like, error)
	DeleteLike(id uuid.UUID, userid uuid.UUID, role string) error
}

type likeService struct {
	Repo repository.LikeRepo
}

func InitLikeService(Repo repository.LikeRepo) LikeService {
	return &likeService{Repo}
}

func (likeService likeService) InsertLike(res dto.LikeRequest) error {
	return likeService.Repo.InsertLike(res)
}

func (likeService likeService) GetLike(page int, limit int, offset int, like string, userid uuid.UUID, blogid uuid.UUID) ([]models.Like, *dto.Pagination, error) {
	return likeService.Repo.GetLike(page, limit, offset, like, userid, blogid)
}

func (likeService likeService) SelectLike(id uuid.UUID) (models.Like, error) {
	return likeService.Repo.SelectLike(id)
}

func (likeService likeService) DeleteLike(id uuid.UUID, userid uuid.UUID, role string) error {
	UserID, err := likeService.Repo.GetLikeUserID(id)
	if err != nil {
		return err
	}

	if UserID != userid && role != "Admin" {
		return errors.New("Access Denied")
	}
	return likeService.Repo.DeleteLike(id)
}

