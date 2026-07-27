package service

import (
	"blog_post/internals/dto"
	"blog_post/internals/repository"
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
)

type ReplyService interface {
	InsertReply(res dto.ReplyRequest) error
	GetReply(page int, limit int, offset int, reply string, blogid uuid.UUID, userid uuid.UUID) ([]models.Reply, *dto.Pagination, error)
	SelectReply(id uuid.UUID) (models.Reply, error)
	UpdateReply(res dto.ReplyRequest, id uuid.UUID) error
	DeleteReply(id uuid.UUID) error
}

type replyService struct {
	Repo repository.ReplyRepo
}

func InitReplyService(Repo repository.ReplyRepo) ReplyService {
	return &replyService{Repo}
}

func (replyService replyService) InsertReply(res dto.ReplyRequest) error {
	return replyService.Repo.InsertReply(res)
}

func (replyService replyService) GetReply(page int, limit int, offset int, reply string, blogid uuid.UUID, userid uuid.UUID) ([]models.Reply, *dto.Pagination, error) {
	return replyService.Repo.GetReply(page, limit, offset, reply, blogid, userid)
}

func (replyService replyService) SelectReply(id uuid.UUID) (models.Reply, error) {
	return replyService.Repo.SelectReply(id)
}

func (replyService replyService) UpdateReply(res dto.ReplyRequest, id uuid.UUID) error { 
	return replyService.Repo.UpdateReply(res, id)
}

func (replyService replyService) DeleteReply(id uuid.UUID) error {
	return replyService.Repo.DeleteReply(id)
}
