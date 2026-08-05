package repository

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	InsertCategory(res dto.CategoryRequest) (models.Category, error)
	GetCategory(page int, limit int, offset int, name string) ([]models.Category, int, error)
	SelectCategory(id uuid.UUID) (models.Category, error)
	UpdateCategory(res dto.CategoryRequest, id uuid.UUID) error
	DeleteCategory(id uuid.UUID) error
}

type categoryRepo struct {
	Db *gorm.DB
}

func InitCategoryRepo(Db *gorm.DB) CategoryRepo {
	return &categoryRepo{Db}
}

func (category categoryRepo) InsertCategory(res dto.CategoryRequest) (models.Category, error) {
	CategoryId, err := uuid.NewV7()
	if err != nil {
		return models.Category{},err
	}

	row := models.Category{
		ID:           CategoryId,
		CategoryName: res.CategoryName,
		Description:  res.Description,
	}

	result := category.Db.Create(&row)
	err = result.Error
	if err != nil {
		return models.Category{},err
	}
	return row,nil
}

func (category categoryRepo) GetCategory(page int, limit int, offset int, name string) ([]models.Category, int, error) {

	var categorys []models.Category

	var count int64

	query := category.Db.Model(&categorys)

	err := query.Count(&count).Error
	if err != nil {
		return []models.Category{}, 0, err
	}

	if name != "" {
		records := query.Where("category_name ILIKE ?", "%"+name+"%").Session(&gorm.Session{})
		if records.Error != nil {
			return []models.Category{}, 0, records.Error
		}
	}

	result := query.Limit(limit).Offset(offset).Find(&categorys)

	if result.RowsAffected == 0 {
		return []models.Category{}, 0, errors.New("Category Record data not found")
	}
	return categorys, int(count), nil
}

func (category categoryRepo) SelectCategory(id uuid.UUID) (models.Category, error) {

	var categorys models.Category

	result := category.Db.First(&categorys, "id =?", id)
	if result.RowsAffected == 0 {
		return models.Category{}, errors.New("Category Record data not found")
	}
	return categorys, nil
}

func (category categoryRepo) UpdateCategory(res dto.CategoryRequest, id uuid.UUID) error {

	var categorys models.Category

	result := category.Db.Model(&categorys).Where("id = ?", id).Updates(models.Category{
		CategoryName: res.CategoryName,
		Description:  res.Description,
	})

	if result.RowsAffected == 0 {
		return errors.New("Category Record data not found")
	}
	return nil
}

func (category categoryRepo) DeleteCategory(id uuid.UUID) error {

	var categorys models.Category
	result := category.Db.Model(&categorys).Delete(&categorys, id)
	if result.RowsAffected == 0 {
		return errors.New("Category Record data not found")
	}
	return nil
}
