package servicemanager

import (
	"../../models/comments"
	"../../models/posts"
	"gorm.io/gorm"
)

type ServiceManager struct {
	CommentService comments.Storage
	PostService    posts.Storage
}

func NewServiceManager(db *gorm.DB) *ServiceManager {
	comRep := comments.NewCommentRepository(db)
	postRep := posts.NewPostRepository(db)

	return &ServiceManager{
		CommentService: comments.NewService(comRep),
		PostService:    posts.NewService(postRep),
	}

}
