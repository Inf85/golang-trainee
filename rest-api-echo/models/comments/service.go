package comments

import (
	"context"
	"gorm.io/gorm"
)

/*
CommentService - Struct
*/
type CommentService struct {
	commentRepository commentRepository
}

/*
NewService - Comment Service Constructor
*/
func NewService(rep commentRepository) Storage {
	return &CommentService{
		commentRepository: rep,
	}
}

/*
GetByID - Get Comment By ID
*/
func (cs *CommentService) GetByID(ctx context.Context, comment []Comments, id string) ([]Comments, error) {
	return cs.commentRepository.GetByID(ctx, comment, id)
}

/*
GetAll = Get All Comments
*/
func (cs *CommentService) GetAll(ctx context.Context, comments []Comments) ([]Comments, error) {
	return cs.commentRepository.GetAll(ctx, comments)
}

/*
DeleteByID - Delete Comment By ID
*/
func (cs *CommentService) DeleteByID(ctx context.Context, comment *Comments, id string) (*gorm.DB, error) {
	return cs.commentRepository.DeleteByID(ctx, comment, id)
}

/*
CreateComment - Create New Comment
*/
func (cs *CommentService) CreateComment(ctx context.Context, comment *Comments) *gorm.DB {
	return cs.commentRepository.CreateComment(ctx, comment)
}
