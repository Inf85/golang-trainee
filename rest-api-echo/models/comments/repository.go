package comments

import (
	"context"
	"gorm.io/gorm"
)

type commentRepository struct {
	database *gorm.DB
}

func NewCommentRepository(database *gorm.DB) commentRepository {
	return commentRepository{database: database}
}

func (cr *commentRepository) GetByID(ctx context.Context, comment []Comments, id string) ([]Comments, error) {
	cr.database.Find(&comment, id)
	return comment, nil
}

func (cr *commentRepository) GetAll(ctx context.Context, comments []Comments) ([]Comments, error) {
	cr.database.Find(&comments)
	return comments, nil
}

func (cr *commentRepository) DeleteByID(ctx context.Context, comment *Comments, id string) (*gorm.DB, error) {
	result := cr.database.Delete(&comment, id)
	return result, nil
}

func (cr *commentRepository) CreateComment(ctx context.Context, comment *Comments) *gorm.DB {
	result := cr.database.Create(&comment)
	return result
}
