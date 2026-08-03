package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type ReplyRepo interface {
	InsertReply(res dto.ReplyRequest) error
	GetReply(page int, limit int, offset int, reply string, commentid uuid.UUID, userid uuid.UUID) ([]models.Reply, *dto.Pagination, error)
	SelectReply(id uuid.UUID) (models.Reply, error)
	UpdateReply(res dto.ReplyRequest, id uuid.UUID) error
	DeleteReply(id uuid.UUID) error

	GetReplyUserID(id uuid.UUID) (uuid.UUID, error)
}

type replyRepo struct {
	Db *gorm.DB
}

func InitReplyRepo(Db *gorm.DB) ReplyRepo {
	return &replyRepo{Db}
}

func (replyRepo replyRepo) InsertReply(res dto.ReplyRequest) error {

	ReplyId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	row := models.Reply{
		ID:        ReplyId,
		Reply:     res.Reply,
		CommentID: res.CommentID,
		UserID:  res.UserID,
	}

	result := replyRepo.Db.Create(&row)
	err = result.Error
	if err != nil {
		return err
	}
	return nil
}

func (replyRepo replyRepo) GetReply(page int, limit int, offset int, reply string, commentid uuid.UUID, userid uuid.UUID) ([]models.Reply, *dto.Pagination, error) {

	var replys []models.Reply

	var count int64

	query := replyRepo.Db.Model(&replys)

	err := query.Count(&count).Error
	if err != nil {
		return nil, nil, err
	}

	if reply != "" {
		records := query.Where("reply LIKE ?", "%"+reply+"%").Session(&gorm.Session{})
		if records.Error != nil {
			return nil, nil, records.Error
		}
	}

	if commentid != uuid.Nil {
		record := query.Where("comment_id = ?", commentid).Session(&gorm.Session{})
		if record.Error != nil {
			return nil, nil, record.Error
		}
	}

	if userid != uuid.Nil {
		record := query.Where("user_id = ?", userid).Session(&gorm.Session{})
		if record.Error != nil {
			return nil, nil, record.Error
		}
	}

	result := query.Limit(limit).Offset(offset).Find(&replys)

	if result.RowsAffected == 0 {
		return nil, nil, errors.New("Like Record data not found")
	}
	return replys, &dto.Pagination{Page: page, Limit: limit, Total: int(count)}, nil
}

func (replyRepo replyRepo) SelectReply(id uuid.UUID) (models.Reply, error) {

	var replys models.Reply

	result := replyRepo.Db.First(&replys, "id =?", id)
	if result.RowsAffected == 0 {
		return models.Reply{}, errors.New("Reply Record data not found")
	}
	return replys, nil
}

func (replyRepo replyRepo) UpdateReply(res dto.ReplyRequest, id uuid.UUID) error {

	var replys models.Reply

	result := replyRepo.Db.Model(&replys).Where("id = ?", id).Updates(models.Reply{
		Reply:     res.Reply,
		CommentID: res.CommentID,
		UserID:    res.UserID,
	})

	if result.RowsAffected == 0 {
		return errors.New("Reply Record data not found")
	}
	return nil
}

func (replyRepo replyRepo) DeleteReply(id uuid.UUID) error {

	var replys models.Reply
	result := replyRepo.Db.Model(&replys).Delete(&replys, id)
	if result.RowsAffected == 0 {
		return errors.New("Reply Record data not found")
	}
	return nil
}

func (replyRepo replyRepo) GetReplyUserID(id uuid.UUID) (uuid.UUID, error) {

	var Replys models.Reply

	result := replyRepo.Db.First(&Replys, "id =?", id)
	if result.RowsAffected == 0 {
		return uuid.Nil, errors.New("Reply Record data not found")
	}

	UserId := Replys.UserID

	return UserId, nil
}
