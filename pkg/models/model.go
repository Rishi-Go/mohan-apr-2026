package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type BlogUsers struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	FirstName    string    `json:"firstname" gorm:"size:100"`
	LastName     string    `json:"lastname" gorm:"size:100"`
	UserName     string    `json:"username" gorm:"size:100"`
	PasswordHash string    `json:"password_hash" gorm:"size:100"`
	Email        string    `json:"email"`
	Role         string    `json:"role" validate:"required,oneof= Admin User"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type Category struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	CategoryName string    `json:"category_name" gorm:"size:100"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type Blog struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Title      string    `json:"title" gorm:"size:100"`
	Content    string    `json:"content" gorm:"size:500"`
	CategoryId uuid.UUID `json:"category_id"`
	Category   Category  `json:"-" gorm:"foreignkey:CategoryId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AuthorID   uuid.UUID `json:"author_id"`
	Users   BlogUsers     `json:"-" gorm:"foreignkey:AuthorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Comment   string    `json:"comment" gorm:"size:500"`
	BlogID    uuid.UUID `json:"blog_id"`
	Blog      Blog      ` json:"-" gorm:"foreignkey:BlogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID    uuid.UUID `json:"user_id"`
	Users  BlogUsers     `json:"-" gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Like struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	LikeResponse bool      `json:"like_response"`
	UserID       uuid.UUID `json:"user_id"`
	BlogUser     BlogUsers     `json:"-" gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type Reply struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Reply     string    `json:"reply" gorm:"size:500"`
	BlogID    uuid.UUID `json:"blog_id"`
	Blog      Blog      ` json:"-" gorm:"foreignkey:BlogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID    uuid.UUID `json:"user_id"`
	BlogUser  BlogUsers     `json:"-" gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
