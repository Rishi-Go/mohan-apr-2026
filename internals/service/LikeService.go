package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
)

type LikeService interface {
	InsertLike(res dto.LikeRequest) error
}

type likeService struct {
	Repo repository.LikeRepo
}

func InitLikeService(Repo repository.LikeRepo) LikeService {
	return &likeService{Repo}
}

func (like likeService) InsertLike(res dto.LikeRequest) error {
	return like.Repo.InsertLike(res)
}

// func (like likeService) GetLike(page int, limit int, offset int, name string) ([]models.Like, *dto.Pagination, error) {
// 	return like.Repo.GetLike(page, limit, offset, name)
// }

// func (like likeService) SelectLike(id uuid.UUID) (models.Like, error) {
// 	return  like.Repo.SelectLike(id)
// }
