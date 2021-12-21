package posts

import (
	"context"
	"gorm.io/gorm"
)

type postRepository struct {
	database *gorm.DB
}

func NewPostRepository(database *gorm.DB) *postRepository {
	return &postRepository{database: database}
}

func (pr *postRepository) GetByID(ctx context.Context, post []Posts, id string) ([]Posts, error) {
	pr.database.Preload("Comments").Find(&post, id)
	return post, nil
}

func (pr *postRepository) GetAll(ctx context.Context, posts []Posts) ([]Posts, error) {
	pr.database.Preload("Comments").Find(&posts)
	return posts, nil
}

func (pr *postRepository) DeleteByID(ctx context.Context, post Posts, id string) (*gorm.DB, error) {
	result := pr.database.Delete(&post, id)
	return result, nil
}

func (pr *postRepository) CreatePost(ctx context.Context, post *Posts) *gorm.DB {
	result := pr.database.Create(&post)
	return result
}
