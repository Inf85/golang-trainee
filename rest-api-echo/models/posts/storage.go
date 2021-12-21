package posts

import (
	"context"
	"gorm.io/gorm"
)

type Storage interface {
	GetByID(ctx context.Context, post []Posts, id string) ([]Posts, error)
	GetAll(ctx context.Context, posts []Posts) ([]Posts, error)
	DeleteByID(ctx context.Context, post Posts, id string) (*gorm.DB, error)
	CreatePost(ctx context.Context, post *Posts) *gorm.DB
}
