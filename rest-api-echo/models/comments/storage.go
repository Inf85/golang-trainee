package comments

import (
	"context"
	"gorm.io/gorm"
)

/*
Storage - Comment Storage Interface
*/
type Storage interface {
	GetByID(ctx context.Context, comment []Comments, id string) ([]Comments, error)
	GetAll(ctx context.Context, comments []Comments) ([]Comments, error)
	DeleteByID(ctx context.Context, comment *Comments, id string) (*gorm.DB, error)
	CreateComment(ctx context.Context, comment *Comments) *gorm.DB
}
