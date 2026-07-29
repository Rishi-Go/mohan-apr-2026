package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepo interface {
	InsertUser(res dto.SignUpRequest) error
	GetUser(page int, limit int, offset int, username string, email string) ([]models.BlogUsers, *dto.Pagination, error)
	SelectUser(id uuid.UUID) (models.BlogUsers, error)
	UpdateUser(res dto.SignUpRequest, id uuid.UUID) error
	DeleteUser(id uuid.UUID) error

	LogInUser(res dto.LogInRequest) (models.BlogUsers, error)
}

type authRepo struct {
	Db *gorm.DB
}

func InitAuthRepo(Db *gorm.DB) AuthRepo {
	return &authRepo{Db}
}

func (auth authRepo) InsertUser(res dto.SignUpRequest) error {
	auth_id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(res.Password), 12)
	if err != nil {
		return err
	}

	password_hash := string(passwordHash)

	row := models.BlogUsers{
		ID:           auth_id,
		FirstName:    res.FirstName,
		LastName:     res.LastName,
		UserName:     res.UserName,
		PasswordHash: password_hash,
		Email:        res.Email,
		Role:         res.Role,
	}

	result := auth.Db.Create(&row)
	err = result.Error
	if err != nil {
		return err
	}
	return nil
}

func (auth authRepo) GetUser(page int, limit int, offset int, username string, email string) ([]models.BlogUsers, *dto.Pagination, error) {

	var users []models.BlogUsers

	var count int64

	query := auth.Db.Model(&users)

	err := query.Count(&count).Error
	if err != nil {
		return nil, nil, err
	}

	if username != "" {
		records := query.Where("user_name LIKE ? AND Role != 'Admin'", "%"+username+"%").Session(&gorm.Session{})
		if records.Error != nil {
			return nil, nil, records.Error
		}
	}

	if email != "" {
		records := query.Where("email LIKE ? AND Role != 'Admin'", "%"+email+"%").Session(&gorm.Session{})
		if records.Error != nil {
			return nil, nil, records.Error
		}
	}
	result := query.Where(" Role != 'Admin'").Limit(limit).Offset(offset).Find(&users)

	if result.RowsAffected == 0 {
		return nil, nil, errors.New("Users Record data not found")
	}
	return users, &dto.Pagination{Page: page, Limit: limit, Total: int(count), Offset: offset}, nil
}

func (auth authRepo) SelectUser(id uuid.UUID) (models.BlogUsers, error) {

	var users models.BlogUsers

	result := auth.Db.First(&users, "id =?", id)
	if result.RowsAffected == 0 {
		return models.BlogUsers{}, errors.New("Users Record data not found")
	}
	return users, nil
}

func (auth authRepo) UpdateUser(res dto.SignUpRequest, id uuid.UUID) error {

	var users models.BlogUsers

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(res.Password), 12)
	if err != nil {
		return err
	}

	password_hash := string(passwordHash)

	result := auth.Db.Model(&users).Where("id = ?", id).Updates(models.BlogUsers{
		FirstName:    res.FirstName,
		LastName:     res.LastName,
		UserName:     res.UserName,
		PasswordHash: password_hash,
		Email:        res.Email,
	})

	if result.RowsAffected == 0 {
		return errors.New("User Record data not found")
	}
	return nil
}

func (auth authRepo) DeleteUser(id uuid.UUID) error {
	var users models.BlogUsers
	result := auth.Db.Model(&users).Delete(&users, id)
	if result.RowsAffected == 0 {
		return errors.New("Users Record data not found")
	}
	return nil
}

func (auth authRepo) LogInUser(res dto.LogInRequest) (models.BlogUsers, error) {

	var users models.BlogUsers

	result := auth.Db.First(&users, "user_name = ?", res.UserName)
	if result.RowsAffected == 0 {
		return models.BlogUsers{}, errors.New("Invalid UserName or Password")
	}

	return users, nil
}
