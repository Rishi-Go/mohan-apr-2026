package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type BlogUsers struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	FirstName    string         `json:"firstname" gorm:"size:100"`
	LastName     string         `json:"lastname" gorm:"size:100"`
	UserName     string         `json:"username" gorm:"size:100,unique"`
	PasswordHash string         `json:"-" gorm:"size:100,unique"`
	Email        string         `json:"email"`
	Role         string         `json:"role" validate:"required,oneof= Admin User" gorm:"default:'User'"`
	CreatedAt    time.Time      `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Category struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	CategoryName string         `json:"category_name" gorm:"size:100"`
	Description  string         `json:"description" gorm:"size:500"`
	CreatedAt    time.Time      `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Blog struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Title      string         `json:"title" gorm:"size:100"`
	Content    string         `json:"content" gorm:"size:500"`
	CategoryId uuid.UUID      `json:"category_id"`
	Category   Category       `json:"-" gorm:"foreignkey:CategoryId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AuthorID   uuid.UUID      `json:"author_id"`
	Users      BlogUsers      `json:"-" gorm:"foreignkey:AuthorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt  time.Time      `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type Comment struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Comment   string         `json:"comment" gorm:"size:500"`
	BlogID    uuid.UUID      `json:"blog_id"`
	Blog      Blog           `json:"-" gorm:"foreignkey:BlogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID    uuid.UUID      `json:"user_id"`
	Users     BlogUsers      `json:"-" gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Like struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	LikeResponse string         `json:"like_response" validate:"required,oneof= Yes No"`
	BlogID       uuid.UUID      `json:"blog_id"`
	Blog         Blog           `json:"-" gorm:"foreignkey:BlogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID       uuid.UUID      `json:"user_id"`
	BlogUser     BlogUsers      `json:"-" gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt    time.Time      `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Reply struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Reply     string         `json:"reply" gorm:"size:500"`
	CommentID uuid.UUID      `json:"comment_id"`
	Comment   Comment        ` json:"-" gorm:"foreignkey:CommentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID    uuid.UUID      `json:"user_id"`
	BlogUser  BlogUsers      `json:"-" gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
