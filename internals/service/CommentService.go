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

type CommentService interface {
	InsertComment(res dto.CommentRequest) (models.Comment, error)
	GetComment(page int, limit int, offset int, comment string, userid uuid.UUID, blogid uuid.UUID) ([]models.Comment, *dto.Pagination, error)
	SelectComment(id uuid.UUID) (models.Comment, error)
	UpdateComment(res dto.CommentRequest, id uuid.UUID) error
	DeleteComment(id uuid.UUID, userid uuid.UUID, role string) error
}

type commentService struct {
	Repo repository.CommentRepo
}

func InitCommentService(Repo repository.CommentRepo) CommentService {
	return &commentService{Repo}
}

func (commentService commentService) InsertComment(res dto.CommentRequest) (models.Comment, error) {
	return commentService.Repo.InsertComment(res)
}

func (commentService commentService) GetComment(page int, limit int, offset int, comment string, userid uuid.UUID, blogid uuid.UUID) ([]models.Comment, *dto.Pagination, error) {

	result, count, err := commentService.Repo.GetComment(page, limit, offset, comment, userid, blogid)
	if err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return []models.Comment{}, nil, err
	}
	return result, &dto.Pagination{Page: page, Limit: limit, Total: int(count)}, err
}

func (commentService commentService) SelectComment(id uuid.UUID) (models.Comment, error) {
	return commentService.Repo.SelectComment(id)
}

func (commentService commentService) UpdateComment(res dto.CommentRequest, id uuid.UUID) error {

	UserID, err := commentService.Repo.GetCommentUserID(id)
	if err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return err
	}

	if UserID != res.UserID {
		logger.Log.Error("Access Denied only comment user")
		return errors.New("Access Denied only comment user")
	}
	return commentService.Repo.UpdateComment(res, id)
}

func (commentService commentService) DeleteComment(id uuid.UUID, userid uuid.UUID, role string) error {

	UserID, err := commentService.Repo.GetCommentUserID(id)
	if err != nil {
		return err
	}

	if UserID != userid && role != "Admin" {
		return errors.New("Access Denied only admin & comment user")
	}
	return commentService.Repo.DeleteComment(id, userid)
}
