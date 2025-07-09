package service

import (
	"gorm.io/gorm"
	"main/model"
	"main/repository"
)

type PostService interface {
	// CreatePost 定义新增文章的方法
	CreatePost(post *model.Posts)
	// GetPost 定义获取单篇文章的方法
	GetPost(id uint) *model.Posts
	// UpdatePost 定义更新文章的方法
	UpdatePost(post *model.Posts)
	// DeletePost 定义删除文章的方法
	DeletePost(id uint)
}

type postService struct {
	db   *gorm.DB
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository, db *gorm.DB) PostService {
	return &postService{
		db:   db,
		repo: repo,
	}
}

func (s *postService) CreatePost(post *model.Posts) {
	s.repo.CreatePost(post)
}

func (s *postService) GetPost(id uint) *model.Posts {
	return s.repo.GetPost(id)
}

func (s *postService) UpdatePost(post *model.Posts) {
	s.repo.UpdatePost(post)
}

func (s *postService) DeletePost(id uint) {
	s.repo.DeletePost(id)
}
