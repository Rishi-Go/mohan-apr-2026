package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
)

type ReplyService interface {
	InsertReply(res dto.ReplyRequest) (models.Reply, error)
	GetReply(page int, limit int, offset int, reply string, commentid uuid.UUID, userid uuid.UUID) ([]models.Reply, *dto.Pagination, error)
	SelectReply(id uuid.UUID) (models.Reply, error)
	UpdateReply(res dto.ReplyRequest, id uuid.UUID) error
	DeleteReply(id uuid.UUID, userid uuid.UUID, role string) error
}

type replyService struct {
	Repo repository.ReplyRepo
}

func InitReplyService(Repo repository.ReplyRepo) ReplyService {
	return &replyService{Repo}
}

func (replyService replyService) InsertReply(res dto.ReplyRequest) (models.Reply, error) {
	return replyService.Repo.InsertReply(res)
}

func (replyService replyService) GetReply(page int, limit int, offset int, reply string, commentid uuid.UUID, userid uuid.UUID) ([]models.Reply, *dto.Pagination, error) {
	result, count, err := replyService.Repo.GetReply(page, limit, offset, reply, commentid, userid)
	if err != nil {
		return []models.Reply{}, nil, err 
	}
	return result, &dto.Pagination{Page: page, Limit: limit, Total: int(count)}, err
}

func (replyService replyService) SelectReply(id uuid.UUID) (models.Reply, error) {
	return replyService.Repo.SelectReply(id)
}

func (replyService replyService) UpdateReply(res dto.ReplyRequest, id uuid.UUID) error {

	UserID, err := replyService.Repo.GetReplyUserID(id)
	if err != nil {
		return err
	}

	if UserID != res.UserID {
		return errors.New("Access Denied only reply user can update")
	}
	return replyService.Repo.UpdateReply(res, id)
}

func (replyService replyService) DeleteReply(id uuid.UUID, userid uuid.UUID, role string) error {

	UserID, err := replyService.Repo.GetReplyUserID(id)
	if err != nil {
		return err
	}

	if UserID != userid && role != "Admin" {
		return errors.New("Access Denied only admin or reply user")
	}

	return replyService.Repo.DeleteReply(id)
}
