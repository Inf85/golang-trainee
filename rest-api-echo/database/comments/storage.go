package comments

import (
	"../../models/comments"
	"gorm.io/gorm"
)

type CommentsStorage struct {
}

func NewStorage() *CommentsStorage {
	return &CommentsStorage{}
}

func (cs *CommentsStorage) GetByID(id string) (*comments.Comments, error) {
	return nil, nil
}

func (cs *CommentsStorage) GetAll() (*comments.Comments, error) {
	return nil, nil
}

func (cs *CommentsStorage) DeleteByID(id string) (*gorm.DB, error) {
	return nil, nil
}
