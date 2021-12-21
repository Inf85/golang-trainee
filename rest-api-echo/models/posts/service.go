package posts

import (
	"context"
	"gorm.io/gorm"
)

type PostService struct {
	postRepository postRepository
}

func NewService(rep *postRepository) Storage {
	return &PostService{
		postRepository: *rep,
	}
}

func (ps *PostService) GetByID(ctx context.Context, post []Posts, id string) ([]Posts, error) {
	return ps.postRepository.GetByID(ctx, post, id)
}

func (ps *PostService) GetAll(ctx context.Context, posts []Posts) ([]Posts, error) {
	return ps.postRepository.GetAll(ctx, posts)
}

func (ps *PostService) DeleteByID(ctx context.Context, post Posts, id string) (*gorm.DB, error) {
	return ps.postRepository.DeleteByID(ctx, post, id)
}

func (ps *PostService) CreatePost(ctx context.Context, post *Posts) *gorm.DB {
	return ps.postRepository.CreatePost(ctx, post)
}
